package indexer

import (
	"fmt"
	"strings"
	"time"

	"github.com/burnfeed/indexer/cmd/flags"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// Config contains all configurations which will be needed by a new action indexer instance.
type Config struct {
	Timeout             time.Duration
	BurnFeedAddress     common.Address
	RPCWSEndpoint       string
	IPFSEndpoint        string
	IPFSProjectID       string
	IPFSProjectSecret   string
	MySqlDsn            string
	IPFSObjectSizeLimit uint64
}

// NewConfigFromCliContext initializes a Config instance from
// command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	burnFeedAddress := c.String(flags.BurnFeedAddress.Name)
	if !common.IsHexAddress(burnFeedAddress) {
		return nil, fmt.Errorf("invalid BurnFeed contract address: %s", burnFeedAddress)
	}

	endpoint := c.String(flags.RPCWSEndpoint.Name)
	if !strings.HasPrefix(endpoint, "ws") {
		return nil, fmt.Errorf("invalid websocket endpoint: %s", endpoint)
	}

	return &Config{
		Timeout:             time.Duration(c.Uint64(flags.RPCTimeout.Name)) * time.Second,
		BurnFeedAddress:     common.HexToAddress(burnFeedAddress),
		RPCWSEndpoint:       endpoint,
		IPFSEndpoint:        c.String(flags.IPFSEndpoint.Name),
		IPFSProjectID:       c.String(flags.IPFSProjectID.Name),
		IPFSProjectSecret:   c.String(flags.IPFSProjectSecret.Name),
		MySqlDsn:            c.String(flags.MySqlDsn.Name),
		IPFSObjectSizeLimit: c.Uint64(flags.IPFSObjectSizeLimit.Name),
	}, nil
}
