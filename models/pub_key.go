package models

import (
	"github.com/burnfeed/indexer/bindings"
	"gorm.io/gorm"
)

type PubKey struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`

	// Block information
	BlockNumber uint64 `gorm:"index"`
	BlockHash   string `gorm:"index"`

	// Shared fields
	User   string `gorm:"index"`
	PubKey []byte
}

func CreatePubKey(db *gorm.DB, e *bindings.BurnFeedProtocolClientPubKey) error {
	result := db.Create(&PubKey{
		BlockNumber: e.Raw.BlockNumber,
		BlockHash:   e.Raw.BlockHash.Hex(),
		User:        e.User.Hex(),
		PubKey:      e.Pubkey,
	})

	return result.Error
}
