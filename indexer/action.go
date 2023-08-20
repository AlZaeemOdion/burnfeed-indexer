package indexer

import (
	"github.com/burnfeed/indexer/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
)

type Action interface {
	Subtype() string
	Burn() uint64
	SaveToDB(db *gorm.DB, header *types.Header, burn uint64) error
}

// AggregatedAction represents one aggregated action.
type AggregatedAction struct {
	Subtype *string `json:"subtype,omitempty"`
	// Used by subtype: tweet / like
	Tweet *string `json:"tweet,omitempty"`
	// Used by subtype: tweet
	RetweetOf *string `json:"retweetOf,omitempty"`
	// Used by subtype: follow
	User     *string `json:"user,omitempty"`
	Followee *string `json:"followee,omitempty"`
	// Used by subtype: send_message
	To      *string `json:"to,omitempty"`
	Message *string `json:"message,omitempty"`
}

// AggregatedActions represents an aggregated actions list which is saved on IPFS.
type AggregatedActions struct {
	Type    string             `json:"type"`
	Version string             `json:"string"`
	Actions []AggregatedAction `json:"actions"`
}

// ToSubActions parses the actions list.
func (a *AggregatedActions) ToSubActions(burn uint64) []Action {
	var actions []Action

	for _, action := range a.Actions {
		if action.Subtype == nil {
			continue
		}

		switch *action.Subtype {
		case models.SubTypeTweet:
			if action.Tweet == nil {
				log.Info("Skip invalid action item", "subtype", action.Subtype)
				continue
			}
			actions = append(actions, &TweetAction{
				ActionCommon: ActionCommon{*action.Subtype, burn},
				Tweet:        *action.Tweet,
				RetweetOf:    action.RetweetOf,
			})
		case models.SubTypeFollow:
			if action.User == nil || action.Followee == nil {
				log.Info("Skip invalid action item", "subtype", action.Subtype)
				continue
			}
			actions = append(actions, &FollowAction{
				ActionCommon: ActionCommon{*action.Subtype, burn},
				User:         common.HexToAddress(*action.User),
				Followee:     common.HexToAddress(*action.Followee),
			})
		case models.SubTypeSendMessage:
			if action.To == nil || action.Message == nil {
				log.Info("Skip invalid action item", "subtype", action.Subtype)
				continue
			}
			actions = append(actions, &SendMessageAction{
				ActionCommon: ActionCommon{*action.Subtype, burn},
				To:           common.HexToAddress(*action.To),
				Message:      *action.Message,
			})
		case models.SubTypeLike:
			if action.Tweet == nil {
				log.Info("Skip invalid action item", "subtype", action.Subtype)
				continue
			}
			actions = append(actions, &LikeAction{
				ActionCommon: ActionCommon{*action.Subtype, burn},
				Tweet:        *action.Tweet,
			})
		default:
			log.Info("Unknown subtype", "subtype", action.Subtype)
		}
	}

	return actions
}

// ActionCommon represents all shared fields for all type of actions.
type ActionCommon struct {
	subtype string
	burn    uint64
}

// Subtype implements Action interface.
func (ac *ActionCommon) Subtype() string { return ac.subtype }

// Burn implements Action interface.
func (ac *ActionCommon) Burn() uint64 { return ac.burn }

// TweetAction represents a "tweet" action.
type TweetAction struct {
	ActionCommon
	Tweet     string
	RetweetOf *string
}

// SaveToDB implements Action interface.
func (t *TweetAction) SaveToDB(db *gorm.DB, header *types.Header, burn uint64) error {
	tweet := &models.Tweet{
		SharedField: models.SharedField{
			BlockNumber: header.Number.Uint64(),
			BlockHash:   header.Hash().Hex(),
			SubType:     models.SubTypeTweet,
			Burn:        burn,
		},
		Tweet: t.Tweet,
	}
	if t.RetweetOf != nil {
		tweet.RetweetOf = t.RetweetOf
	}

	result := db.Create(tweet)

	return result.Error
}

// FollowAction represents a "follow" action.
type FollowAction struct {
	ActionCommon
	User     common.Address
	Followee common.Address
}

// SaveToDB implements Action interface.
func (f *FollowAction) SaveToDB(db *gorm.DB, header *types.Header, burn uint64) error {
	result := db.Create(&models.Follow{
		SharedField: models.SharedField{
			BlockNumber: header.Number.Uint64(),
			BlockHash:   header.Hash().Hex(),
			SubType:     models.SubTypeTweet,
			Burn:        burn,
		},
		User:     f.User.Hex(),
		Followee: f.Followee.Hex(),
	})

	return result.Error
}

// LikeAction represents a "like" action.
type LikeAction struct {
	ActionCommon
	Tweet string
}

// SaveToDB implements Action interface.
func (l *LikeAction) SaveToDB(db *gorm.DB, header *types.Header, burn uint64) error {
	result := db.Create(&models.Like{
		SharedField: models.SharedField{
			BlockNumber: header.Number.Uint64(),
			BlockHash:   header.Hash().Hex(),
			SubType:     models.SubTypeTweet,
			Burn:        burn,
		},
		Tweet: l.Tweet,
	})

	return result.Error
}

// SendMessageAction represents a "send_message" action.
type SendMessageAction struct {
	ActionCommon
	To      common.Address
	Message string
}

// SaveToDB implements Action interface.
func (s *SendMessageAction) SaveToDB(db *gorm.DB, header *types.Header, burn uint64) error {
	result := db.Create(&models.SendMessage{
		SharedField: models.SharedField{
			BlockNumber: header.Number.Uint64(),
			BlockHash:   header.Hash().Hex(),
			SubType:     models.SubTypeTweet,
			Burn:        burn,
		},
		To:      s.To.Hex(),
		Message: s.Message,
	})

	return result.Error
}
