# burnfeed-indexer
[![CI](https://github.com/AlZaeemOdion/burnfeed-indexer/actions/workflows/test.yml/badge.svg)](https://github.com/AlZaeemOdion/burnfeed-indexer/actions/workflows/test.yml)

An indexer implementation for burnFeed protocol

## Usage

Start a local development environment, including:
- a foundry node (port: 8545)
- a MySQL instance (port: 3306)
- an IPFS node (port: 5001 / webui: http://localhost:5001/webui)
- a burnFeed indexer

```sh
BURN_FEED_DIR=<BURN_FEED_PROTOCOL_DIR> \
  make dev
```

Run tests:
```sh
BURN_FEED_DIR=<BURN_FEED_PROTOCOL_DIR> \
  make test
```

## Tables

See [here](https://github.com/AlZaeemOdion/burnfeed-indexer/blob/main/scripts/ddl.sql).
