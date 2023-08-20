#!/bin/bash

# Generate go contract bindings.
# ref: https://geth.ethereum.org/docs/dapp/native-bindings

set -eou pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

echo ""
echo "BURN_FEED_DIR: ${BURN_FEED_DIR}"
echo "GETH_DIR: ${GETH_DIR}"
echo ""

cd ${GETH_DIR} &&
  make all &&
  cd -

cd ${BURN_FEED_DIR} &&
  forge install &&
  forge compile &&
  cd -

ABIGEN_BIN=$GETH_DIR/build/bin/abigen

echo ""
echo "Start generating go contract bindings..."
echo ""

cat ${BURN_FEED_DIR}/out/BurnFeedProtocol.sol/BurnFeedProtocol.json |
	jq .abi |
	${ABIGEN_BIN} --abi - --type BurnFeedProtocolClient --pkg bindings --out $DIR/../bindings/gen_burn_feed.go

echo "🍻 Go contract bindings generated!"
