# burnfeed-indexer
[![CI](https://github.com/AlZaeemOdion/burnfeed-indexer/actions/workflows/test.yml/badge.svg)](https://github.com/AlZaeemOdion/burnfeed-indexer/actions/workflows/test.yml)

An indexer implementation for burnFeed protocol

## Usage

Start a local development environment (if `BURN_FEED_DIR` environment variable is not specified, will use `./burnfeed_home` as default value), including:
- a foundry node (port: 8545)
- a MySQL instance (port: 3306)
- an IPFS node (port: 5001 / webui: http://localhost:5001/webui)
- a burnFeed indexer

```sh
BURN_FEED_DIR=<BURN_FEED_PROTOCOL_DIR> make dev
```

Run tests:
```sh
BURN_FEED_DIR=<BURN_FEED_PROTOCOL_DIR> make test
```

## Flags

```
GLOBAL OPTIONS:
   --rpc.wsEndpoint value    Websocket RPC endpoint of an Ethereum node
   --ipfs.endpoint value     An IPFS API endpoint
   --ipfs.apiKey value       An IPFS API key
   --ipfs.apiSecret value    An IPFS API secret
   --burnFeed.address value  BurnFeed contract address
   --mysql.dsn value         DSN of the MySQL database, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
   --rpc.timeout value       Timeout for the RPC requests in second (default: 60)
   --ipfs.sizeLimit value    Size limit of the ipfs object, if an object exceeds this limit, then it won't be indexed (default: 10240)
   --verbosity value         Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail (default: 3)
   --log.json                Format logs with JSON (default: false)
   --help, -h                show help
```

## Tables

See [here](https://github.com/AlZaeemOdion/burnfeed-indexer/blob/main/scripts/ddl.sql).
