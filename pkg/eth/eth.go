package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ntc-goer/parser-exercise/config"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"time"
)

type ETH struct {
	Client *ethclient.Client
}

func NewETH(cfg *config.Config) *ETH {
	client, err := ethclient.Dial(cfg.InfuraUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return &ETH{
		Client: client,
	}
}

func (e *ETH) GetLatestBlockNumber() (int64, error) {
	ctx := context.TODO()
	curBlock, err := e.Client.BlockByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return curBlock.Number().Int64(), nil
}

func (e *ETH) GetAllCheckBlockSinceTime(subscribeTime time.Time, fromBlock int64) ([]*types.Block, error) {
	ctx := context.TODO()
	latestBlockNumber, err := e.GetLatestBlockNumber()
	if err != nil {
		return nil, err
	}
	blocks := make([]*types.Block, 0)
	for i := latestBlockNumber; i >= fromBlock; i-- {
		block, err := e.Client.BlockByNumber(ctx, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		blockTime := time.Unix(int64(block.Time()), 0).UTC()
		blocks = append(blocks, block)
		if blockTime.Before(subscribeTime) {
			break
		}
	}
	return blocks, nil
}

func (e *ETH) FilterTransactionFromBlock(block *types.Block, condition func(tx *types.Transaction) bool) ([]*types.Transaction, error) {
	validTransactions := make([]*types.Transaction, 0)
	for _, tx := range block.Transactions() {
		if tx == nil {
			logrus.Info(tx)
		}
		if condition(tx) {
			validTransactions = append(validTransactions, tx)
		}
	}
	return validTransactions, nil
}
