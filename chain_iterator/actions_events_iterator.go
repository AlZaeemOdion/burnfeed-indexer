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

// EndActionsEventIterFunc ends the current iteration.
type EndActionsEventIterFunc func()

// OnActionsEvent represents the callback function which will be called when a Burnfeed.Actions event is
// iterated.
type OnActionsEvent func(context.Context, *bindings.BurnFeedProtocolClientActions, EndActionsEventIterFunc) error

// ActionsIterator iterates the emitted Burnfeed.Actions events in the chain.
type ActionsIterator struct {
	ctx                context.Context
	burnFeed           *bindings.BurnFeedProtocolClient
	blockBatchIterator *chainIterator.BlockBatchIterator
	isEnd              bool
}

// ActionsIteratorConfig represents the configs of a Actions event iterator.
type ActionsIteratorConfig struct {
	Client                *rpc.EthClient
	BurnFeed              *bindings.BurnFeedProtocolClient
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	OnActionsEvent        OnActionsEvent
}

// NewActionsIterator creates a new instance of Actions event iterator.
func NewActionsIterator(ctx context.Context, cfg *ActionsIteratorConfig) (*ActionsIterator, error) {
	if cfg.OnActionsEvent == nil {
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
		OnBlocks: assembleActionsIteratorCallback(
			cfg.Client,
			cfg.BurnFeed,
			cfg.OnActionsEvent,
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
func (i *ActionsIterator) Iter() error {
	return i.blockBatchIterator.Iter()
}

// end ends the current iteration.
func (i *ActionsIterator) end() {
	i.isEnd = true
}

// assembleActionsIteratorCallback assembles the callback which will be used
// by a event iterator's inner block iterator.
func assembleActionsIteratorCallback(
	client *rpc.EthClient,
	burnFeed *bindings.BurnFeedProtocolClient,
	callback OnActionsEvent,
	eventIter *ActionsIterator,
) chainIterator.OnBlocksFunc {
	return func(
		ctx context.Context,
		start, end *types.Header,
		updateCurrentFunc chainIterator.UpdateCurrentFunc,
		endFunc chainIterator.EndIterFunc,
	) error {
		endHeight := end.Number.Uint64()
		iter, err := burnFeed.FilterActions(
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
