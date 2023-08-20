package models

import (
	"context"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// Valid subtypes.
var (
	SubTypeTweet       = "tweet"
	SubTypeFollow      = "follow"
	SubTypeSendMessage = "send_message"
	SubTypeLike        = "like"
)

// AutoMigrate runs auto migration for given models.
func AutoMigrate(ctx context.Context, db *gorm.DB) error {
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error { return db.AutoMigrate(new(Tweet)) })
	g.Go(func() error { return db.AutoMigrate(new(Follow)) })
	g.Go(func() error { return db.AutoMigrate(new(SendMessage)) })
	g.Go(func() error { return db.AutoMigrate(new(Like)) })
	g.Go(func() error { return db.AutoMigrate(new(PubKey)) })

	return g.Wait()
}
