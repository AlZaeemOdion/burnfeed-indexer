package flags

import (
	"github.com/urfave/cli/v2"
)

// Required flags.
var (
	RPCWSEndpoint = &cli.StringFlag{
		Name:     "rpc.wsEndpoint",
		Usage:    "Websocket RPC endpoint of an Ethereum node",
		Required: true,
	}
	IPFSEndpoint = &cli.StringFlag{
		Name:     "ipfs.endpoint",
		Usage:    "An IPFS API endpoint",
		Required: true,
	}
	IPFSProjectID = &cli.StringFlag{
		Name:     "ipfs.apiKey",
		Usage:    "An IPFS API key",
		Required: true,
	}
	IPFSProjectSecret = &cli.StringFlag{
		Name:     "ipfs.apiSecret",
		Usage:    "An IPFS API secret",
		Required: true,
	}
	BurnFeedAddress = &cli.StringFlag{
		Name:     "burnFeed.address",
		Usage:    "BurnFeed contract address",
		Required: true,
	}
	MySqlDsn = &cli.StringFlag{
		Name:     "mysql.dsn",
		Usage:    "DSN of the MySQL database, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details",
		Required: true,
	}
	// Optional flags.
	RPCTimeout = &cli.UintFlag{
		Name:  "rpc.timeout",
		Usage: "Timeout for the RPC requests in second",
		Value: 60,
	}
	IPFSObjectSizeLimit = &cli.UintFlag{
		Name:  "ipfs.sizeLimit",
		Usage: "Size limit of the ipfs object, if an object exceeds this limit, then it won't be indexed",
		Value: 1024, // TODO: update this value
	}
	Verbosity = &cli.IntFlag{
		Name:  "verbosity",
		Usage: "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail",
		Value: 3,
	}
	LogJson = &cli.BoolFlag{
		Name:  "log.json",
		Usage: "Format logs with JSON",
	}
)

// All common flags.
var CommonFlags = []cli.Flag{
	// Required
	RPCWSEndpoint,
	IPFSEndpoint,
	IPFSProjectID,
	IPFSProjectSecret,
	BurnFeedAddress,
	MySqlDsn,
	// Optional
	RPCTimeout,
	IPFSObjectSizeLimit,
	Verbosity,
	LogJson,
}
