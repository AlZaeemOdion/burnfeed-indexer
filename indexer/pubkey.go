package indexer

import (
	"github.com/burnfeed/indexer/bindings"
	"github.com/burnfeed/indexer/models"
	"github.com/ethereum/go-ethereum/core/types"
)

// SavePubKey saves the given PubKey event to database.
func (i *ActionsIndexer) SavePubKey(event *bindings.BurnFeedProtocolClientPubKey, header *types.Header) error {
	result := i.db.Create(&models.PubKey{
		BlockNumber: header.Number.Uint64(),
		BlockHash:   header.Hash().Hex(),
		User:        event.User.Hex(),
		PubKey:      event.Pubkey,
	})

	return result.Error
}
