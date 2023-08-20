package indexer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/burnfeed/indexer/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/suite"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

type IndexerTestSuite struct {
	suite.Suite
	TestAddrPrivKey *ecdsa.PrivateKey
	TestAddr        common.Address
	Indexer         *ActionsIndexer
	ChainID         *big.Int
	ctx             context.Context
	cancel          context.CancelFunc
}

func (s *IndexerTestSuite) SetupTest() {
	// Default logger
	log.Root().SetHandler(
		log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stdout, log.TerminalFormat(true))),
	)

	if os.Getenv("LOG_LEVEL") != "" {
		level, err := log.LvlFromString(os.Getenv("LOG_LEVEL"))
		if err != nil {
			log.Crit("Invalid log level", "level", os.Getenv("LOG_LEVEL"))
		}
		log.Root().SetHandler(
			log.LvlFilterHandler(level, log.StreamHandler(os.Stdout, log.TerminalFormat(true))),
		)
	}

	// Default keys used by foundry node.
	testAddrPrivKey, err := crypto.ToECDSA(common.Hex2Bytes(os.Getenv("TEST_PRIV_KEY")))
	s.Nil(err)
	s.TestAddrPrivKey = testAddrPrivKey
	s.TestAddr = common.HexToAddress(os.Getenv("TEST_ADDR"))

	s.ctx, s.cancel = context.WithCancel(context.Background())
	indexer, err := New(s.ctx, &Config{
		Timeout:             1 * time.Minute,
		BurnFeedAddress:     common.HexToAddress(os.Getenv("BURN_FEED_ADDRESS")),
		RPCWSEndpoint:       os.Getenv("RPC_WS_ENDPOINT"),
		IPFSEndpoint:        os.Getenv("IPFS_ENDPOINT"),
		IPFSProjectID:       os.Getenv("IPFS_API_KEY"),
		IPFSProjectSecret:   os.Getenv("IPFS_API_SECRET"),
		MySqlDsn:            os.Getenv("MYSQL_DSN"),
		IPFSObjectSizeLimit: 1024000,
	})
	s.Nil(err)
	s.Indexer = indexer

	s.ChainID, err = indexer.ethClient.ChainID(context.Background())
	s.Nil(err)
}

func (s *IndexerTestSuite) TestIndexOpPubKey() {
	// Publish a private key
	opts, err := bind.NewKeyedTransactorWithChainID(s.TestAddrPrivKey, s.ChainID)
	s.Nil(err)
	tx, err := s.Indexer.burnFeedClient.RegisterPubKey(opts, []byte{1})
	s.Nil(err)
	_, err = rpc.WaitReceipt(context.Background(), s.Indexer.ethClient, tx)
	s.Nil(err)

	newHead, err := s.Indexer.ethClient.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(newHead.Number.Uint64(), s.Indexer.currentBlockCursor.Number.Uint64())
	s.Greater(newHead.Number.Uint64(), s.Indexer.head.Number.Uint64())
	s.Indexer.updateChainHead(newHead)

	s.Nil(s.Indexer.indexOp())

	var savedPubKey models.PubKey
	result := s.Indexer.db.First(&savedPubKey, nil)
	log.Info(
		"Saved PubKey",
		"ID", savedPubKey.ID,
		"user", savedPubKey.User,
		"PubKey", savedPubKey.PubKey,
		"BlockNumber", savedPubKey.BlockNumber,
		"blockHash", savedPubKey.BlockHash,
	)
	s.Nil(result.Error)
	s.Equal(s.TestAddr.Hex(), savedPubKey.User)
}

func (s *IndexerTestSuite) TestIndexOpActions() {
	var (
		testIpfsUri           = "ipfs:QmZkH64BFAkVVhoFAPA8uBkfNyzmQeKSUqZoGUXPNzXdC9"
		testUserAddr          = s.TestAddr.Hex()
		testAggregatedActions = &AggregatedActions{
			Type:    TypeAggregatedActions,
			Version: time.Now().String(),
			Actions: []AggregatedAction{
				// Tweet
				{
					Subtype: &models.SubTypeTweet,
					Tweet:   &testIpfsUri,
				},
				// Retweet
				{
					Subtype:   &models.SubTypeTweet,
					Tweet:     &testIpfsUri,
					RetweetOf: &testIpfsUri,
				},
				// Follow
				{
					Subtype:  &models.SubTypeFollow,
					User:     &testUserAddr,
					Followee: &testUserAddr,
				},
				// Invalid Follow
				{
					Subtype: &models.SubTypeFollow,
					User:    &testUserAddr,
				},
				// Like
				{
					Subtype: &models.SubTypeLike,
					Tweet:   &testIpfsUri,
				},
				// SendMessage
				{
					Subtype: &models.SubTypeSendMessage,
					To:      &testUserAddr,
					Message: &testIpfsUri,
				},
			},
		}
	)
	// Publish an aggregated actions list
	jsonPayload, err := json.Marshal(&testAggregatedActions)
	s.Nil(err)
	s.NotEmpty(jsonPayload)

	log.Info("Aggregated actions", "payload", string(jsonPayload))

	cid, err := s.Indexer.ipfsClient.Add(bytes.NewReader(jsonPayload))
	s.Nil(err)

	log.Info("CID", "path", cid)

	opts, err := bind.NewKeyedTransactorWithChainID(s.TestAddrPrivKey, s.ChainID)
	s.Nil(err)
	tx, err := s.Indexer.burnFeedClient.PublishActions(opts, "ipfs:"+cid, common.Big0)
	s.Nil(err)
	_, err = rpc.WaitReceipt(context.Background(), s.Indexer.ethClient, tx)
	s.Nil(err)

	newHead, err := s.Indexer.ethClient.HeaderByNumber(context.Background(), nil)
	s.Nil(err)
	s.Greater(newHead.Number.Uint64(), s.Indexer.currentBlockCursor.Number.Uint64())
	s.Greater(newHead.Number.Uint64(), s.Indexer.head.Number.Uint64())
	s.Indexer.updateChainHead(newHead)

	s.Nil(s.Indexer.indexOp())

	var savedTweet models.Tweet
	result := s.Indexer.db.First(&savedTweet, nil)
	log.Info("Saved tweet", "tweet", savedTweet)
	s.Nil(result.Error)
	s.Equal(models.SubTypeTweet, savedTweet.SubType)
}

func TestIndexerTestSuite(t *testing.T) {
	suite.Run(t, new(IndexerTestSuite))
}
