package chain_iterator

import (
	"context"
	"errors"
	"math/big"

	"github.com/burnfeed/indexer/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	chainIterator "github.com/taikoxyz/taiko-client/pkg/chain_iterator"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// EndPubKeyEventIterFunc ends the current iteration.
type EndPubKeyEventIterFunc func()

// OnPubKeyEvent represents the callback function which will be called when a Burnfeed.Actions event is
// iterated.
type OnPubKeyEvent func(context.Context, *bindings.BurnFeedProtocolClientPubKey, EndPubKeyEventIterFunc) error

// PubKeyIterator iterates the emitted Burnfeed.PubKey events in the chain.
type PubKeyIterator struct {
	ctx                context.Context
	burnFeed           *bindings.BurnFeedProtocolClient
	blockBatchIterator *chainIterator.BlockBatchIterator
	isEnd              bool
}

// PubKeyIteratorConfig represents the configs of a PubKey event iterator.
type PubKeyIteratorConfig struct {
	Client                *rpc.EthClient
	BurnFeed              *bindings.BurnFeedProtocolClient
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	OnPubKeyEvent         OnPubKeyEvent
}

// NewPubKeyIterator creates a new instance of PubKey event iterator.
func NewPubKeyIterator(ctx context.Context, cfg *PubKeyIteratorConfig) (*ActionsIterator, error) {
	if cfg.OnPubKeyEvent == nil {
		return nil, errors.New("invalid callback")
	}

	iterator := &ActionsIterator{
		ctx:      ctx,
		burnFeed: cfg.BurnFeed,
	}

	// Initialize the inner block iterator.
	blockIterator, err := chainIterator.NewBlockBatchIterator(ctx, &chainIterator.BlockBatchIteratorConfig{
		Client:                cfg.Client,
		MaxBlocksReadPerEpoch: cfg.MaxBlocksReadPerEpoch,
		StartHeight:           cfg.StartHeight,
		EndHeight:             cfg.EndHeight,
		OnBlocks: assemblePubKeyIteratorCallback(
			cfg.Client,
			cfg.BurnFeed,
			cfg.OnPubKeyEvent,
			iterator,
		),
	})
	if err != nil {
		return nil, err
	}

	iterator.blockBatchIterator = blockIterator

	return iterator, nil
}

// Iter iterates the given chain between the given start and end heights,
// will call the callback when a Actions event is iterated.
func (i *PubKeyIterator) Iter() error {
	return i.blockBatchIterator.Iter()
}

// end ends the current iteration.
func (i *PubKeyIterator) end() {
	i.isEnd = true
}

// assemblePubKeyIteratorCallback assembles the callback which will be used
// by a event iterator's inner block iterator.
func assemblePubKeyIteratorCallback(
	client *rpc.EthClient,
	burnFeed *bindings.BurnFeedProtocolClient,
	callback OnPubKeyEvent,
	eventIter *ActionsIterator,
) chainIterator.OnBlocksFunc {
	return func(
		ctx context.Context,
		start, end *types.Header,
		updateCurrentFunc chainIterator.UpdateCurrentFunc,
		endFunc chainIterator.EndIterFunc,
	) error {
		endHeight := end.Number.Uint64()
		iter, err := burnFeed.FilterPubKey(
			&bind.FilterOpts{Start: start.Number.Uint64(), End: &endHeight, Context: ctx},
			nil,
		)
		if err != nil {
			return err
		}
		defer iter.Close()

		for iter.Next() {
			event := iter.Event

			if err := callback(ctx, event, eventIter.end); err != nil {
				return err
			}

			if eventIter.isEnd {
				endFunc()
				return nil
			}

			current, err := client.HeaderByHash(ctx, event.Raw.BlockHash)
			if err != nil {
				return err
			}

			updateCurrentFunc(current)
		}

		return nil
	}
}
