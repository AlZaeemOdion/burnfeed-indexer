# burnfeed-indexer
[![CI](https://github.com/AlZaeemOdion/burnfeed-indexer/actions/workflows/test.yml/badge.svg)](https://github.com/AlZaeemOdion/burnfeed-indexer/actions/workflows/test.yml)

An indexer implementation for burnFeed protocol

## Usage

Start a local development environment, including:
- a foundry node (port: 8545)
- a MySQL instance (port: 3306)
- a burnFeed indexer

```sh
IPFS_ENDPOINT=<YOUR_IPFS_ENDPOINT> \
IPFS_API_KEY=<YOUR_IPFS_API_KEY> \
IPFS_API_SECRET=<YOUR_IPFS_API_SECRET> \
BURN_FEED_DIR=<BURN_FEED_PROTOCOL_DIR> \
  make dev
```

## Tables

See [here](https://github.com/AlZaeemOdion/burnfeed-indexer/blob/main/scripts/ddl.sql).
