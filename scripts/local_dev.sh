#!/bin/bash

set -eou pipefail

DIR=$(
    cd $(dirname ${BASH_SOURCE[0]})
    pwd
)
COMPOSE_CONFIG=$DIR/docker-compose.yml

if ! command -v docker &>/dev/null 2>&1; then
    echo "ERROR: docker command not found"
    exit 1
fi

if ! docker info >/dev/null 2>&1; then
    echo "ERROR: docker daemon isn't running"
    exit 1
fi

echo "Starting docker-compose network..."

docker compose -f $COMPOSE_CONFIG down -v --remove-orphans &>/dev/null
docker compose -f $COMPOSE_CONFIG up -d

NODE_URL=localhost:8545 $DIR/wait_for_node.sh
BURN_FEED_DIR="${BURN_FEED_DIR:-./burnfeed_home}"

# Deploy burnFeed protocol.
cd $BURN_FEED_DIR &&
  forge build &&
  forge create \
    --rpc-url http://localhost:8545 \
    --constructor-args "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" \
    --private-key ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    src/BurnFeedProtocol.sol:BurnFeedProtocol

trap "docker compose -f $COMPOSE_CONFIG down -v" EXIT INT KILL ERR
if [ "$RUN_TESTS" == "true" ]; then
    cd $DIR/.. && BURN_FEED_ADDRESS=0x5FbDB2315678afecb367f032d93F642f64180aa3 \
    RPC_WS_ENDPOINT=ws://localhost:8546 \
    IPFS_ENDPOINT=http://localhost:8080 \
    TEST_ADDR="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" \
    TEST_PRIV_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    LOG_LEVEL=debug \
    MYSQL_DSN="burnfeed:burnfeed@tcp(localhost:3306)/burnfeed?charset=utf8mb4&parseTime=True&loc=Local" \
        go test -v -p=1 ./... -coverprofile=coverage.out -covermode=atomic -timeout=300s
else
    cd $DIR/.. && go run ./cmd/main.go \
    --rpc.wsEndpoint ws://localhost:8546 \
    --ipfs.endpoint http://localhost:8080 \
    --burnFeed.address "0x5FbDB2315678afecb367f032d93F642f64180aa3" \
    --mysql.dsn "burnfeed:burnfeed@tcp(localhost:3306)/burnfeed?charset=utf8mb4&parseTime=True&loc=Local" \
    --verbosity "4"
fi
