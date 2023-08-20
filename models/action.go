package models

import (
	"gorm.io/gorm"
)

type SharedField struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`

	// Block information
	BlockNumber uint64 `gorm:"index"`
	BlockHash   string `gorm:"index"`

	SubType string `gorm:"index"`
	Burn    uint64 `gorm:"index"`
}

type Tweet struct {
	SharedField
	Tweet     string
	RetweetOf *string
}

type Follow struct {
	SharedField
	User     string `gorm:"index"`
	Followee string
}

type SendMessage struct {
	SharedField
	To      string `gorm:"index"`
	Message string
}

type Like struct {
	SharedField
	Tweet string `gorm:"index"`
}
