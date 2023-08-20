package indexer

import (
	"context"
	"fmt"
	"sync"

	"github.com/burnfeed/indexer/bindings"
	chainIterator "github.com/burnfeed/indexer/chain_iterator"
	"github.com/burnfeed/indexer/models"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	ipfsApi "github.com/ipfs/go-ipfs-api"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	maxBlocksReadPerEpoch uint64 = 100
)

type ActionsIndexer struct {
	ctx        context.Context
	db         *gorm.DB
	ethClient  *rpc.EthClient
	ipfsClient *ipfsApi.Shell

	burnFeedClient     *bindings.BurnFeedProtocolClient
	currentBlockCursor *types.Header

	sizeLimit uint64

	head          *types.Header
	newHeadSub    event.Subscription
	newHeadCh     chan *types.Header
	newHeadNotify chan struct{}
	headMutex     sync.RWMutex
}

// New creates a new ActionsIndexer instance with the given configs.
func New(ctx context.Context, cfg *Config) (*ActionsIndexer, error) {
	client, err := ethclient.DialContext(ctx, cfg.RPCWSEndpoint)
	if err != nil {
		return nil, err
	}
	clientWithTimeout := rpc.NewEthClientWithTimeout(client, cfg.Timeout)

	burnFeedClient, err := bindings.NewBurnFeedProtocolClient(
		cfg.BurnFeedAddress,
		clientWithTimeout,
	)
	if err != nil {
		return nil, err
	}

	head, err := clientWithTimeout.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(cfg.MySqlDsn), nil)
	if err != nil {
		return nil, err
	}

	if err := models.AutoMigrate(ctx, db); err != nil {
		return nil, err
	}

	newHeadCh := make(chan *types.Header, 1024)
	indexer := &ActionsIndexer{
		ctx:            ctx,
		db:             db,
		ethClient:      clientWithTimeout,
		ipfsClient:     ipfsApi.NewShellWithClient(cfg.IPFSEndpoint, NewClient(cfg.IPFSProjectID, cfg.IPFSProjectSecret)),
		burnFeedClient: burnFeedClient,
		sizeLimit:      cfg.IPFSObjectSizeLimit,
		head:           head,
		newHeadCh:      newHeadCh,
		newHeadNotify:  make(chan struct{}, 1),
		newHeadSub:     rpc.SubscribeChainHead(clientWithTimeout, newHeadCh),
	}

	if err := indexer.initCurrentBlockCursor(); err != nil {
		return nil, err
	}

	return indexer, nil
}

// Start makes the given action indexer starting indexing.
func (i *ActionsIndexer) Start() {
	reqIndexing := func() {
		select {
		case i.newHeadNotify <- struct{}{}:
		default:
		}
	}

	reqIndexing()

	for {
		select {
		case <-i.ctx.Done():
			return
		case <-i.newHeadNotify:
			if err := i.indexOp(); err != nil {
				log.Error("Failed to index actions", "error", err)
			}
		case header := <-i.newHeadCh:
			i.updateChainHead(header)
			reqIndexing()
		}
	}
}

func (i *ActionsIndexer) initCurrentBlockCursor() error {
	head, err := i.ethClient.HeaderByNumber(i.ctx, nil)
	if err != nil {
		return err
	}

	log.Info("Init current block cursor", "number", head.Number, "hash", head.Hash())

	i.currentBlockCursor = head
	return nil
}

func (i *ActionsIndexer) updateChainHead(newHead *types.Header) {
	i.headMutex.Lock()
	defer i.headMutex.Unlock()

	log.Info("New chain head", "number", newHead.Number, "hash", newHead.Hash())

	i.head = newHead
}

func (i *ActionsIndexer) readChainHead() *types.Header {
	i.headMutex.RLock()
	defer i.headMutex.RUnlock()

	return i.head
}

func (i *ActionsIndexer) indexOp() error {
	head := i.readChainHead()

	// Index PubKey events at first.
	iterPubKey, err := chainIterator.NewPubKeyIterator(i.ctx, &chainIterator.PubKeyIteratorConfig{
		Client:                i.ethClient,
		BurnFeed:              i.burnFeedClient,
		MaxBlocksReadPerEpoch: &maxBlocksReadPerEpoch,
		StartHeight:           i.currentBlockCursor.Number,
		EndHeight:             head.Number,
		OnPubKeyEvent:         i.onNewPubKey,
	})
	if err != nil {
		return fmt.Errorf("failed to create event iterator: %w", err)
	}
	if err := iterPubKey.Iter(); err != nil {
		return err
	}

	// Index aggregated actions.
	iterActions, err := chainIterator.NewActionsIterator(i.ctx, &chainIterator.ActionsIteratorConfig{
		Client:                i.ethClient,
		BurnFeed:              i.burnFeedClient,
		MaxBlocksReadPerEpoch: &maxBlocksReadPerEpoch,
		StartHeight:           i.currentBlockCursor.Number,
		EndHeight:             head.Number,
		OnActionsEvent:        i.onNewAction,
	})
	if err != nil {
		return fmt.Errorf("failed to create event iterator: %w", err)
	}
	if err := iterActions.Iter(); err != nil {
		return err
	}

	// Update block cursor.
	i.currentBlockCursor = head

	return nil
}

// onNewAction is the Actions event handler of current indexer.
func (i *ActionsIndexer) onNewAction(
	ctx context.Context,
	event *bindings.BurnFeedProtocolClientActions,
	endIter chainIterator.EndActionsEventIterFunc,
) (err error) {
	header, err := i.ethClient.HeaderByHash(i.ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}

	log.Info("New actions", "user", event.User, "uri", event.Uri, "burn", event.Burn)

	actions, err := i.GetActionsByUri(event.Uri, event.Burn.Uint64())
	if err != nil {
		if e, ok := err.(*InvalidActionEventError); ok {
			log.Info(
				"Invalid action event",
				"blockNumber", header.Number,
				"blockHash", header.Hash(),
				"error", e.Error(),
			)
			return nil
		}
		return err
	}

	for _, action := range actions {
		if err := action.SaveToDB(i.db, header, event.Burn.Uint64()); err != nil {
			return err
		}
	}

	return nil
}

// onNewPubKey is the PubKey event handler of current indexer.
func (i *ActionsIndexer) onNewPubKey(
	ctx context.Context,
	event *bindings.BurnFeedProtocolClientPubKey,
	endIter chainIterator.EndPubKeyEventIterFunc,
) (err error) {
	header, err := i.ethClient.HeaderByHash(i.ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}

	log.Info("New PubKey", "user", event.User, "pubKey", event.Pubkey)

	return i.SavePubKey(event, header)
}
