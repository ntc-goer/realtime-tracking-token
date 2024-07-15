package worker

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ntc-goer/parser-exercise/internal/subscribe"
	"github.com/ntc-goer/parser-exercise/internal/transaction"
	"github.com/ntc-goer/parser-exercise/pkg/eth"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"time"
)

type GetTransactionWorker struct {
	ETH                *eth.ETH
	TransactionService *transaction.Service
	SubscribeService   *subscribe.Service
}

func NewGetTransactionWorker(eth *eth.ETH, ts *transaction.Service, ss *subscribe.Service) *GetTransactionWorker {
	return &GetTransactionWorker{
		ETH:                eth,
		TransactionService: ts,
		SubscribeService:   ss,
	}
}

func (wk *GetTransactionWorker) Run(ctx context.Context, sub subscribe.Subscribe) {
	logrus.Infof("Collect transaction of address %s", sub.Address)
	blocks, err := wk.ETH.GetAllCheckBlockSinceTime(sub.CreatedAt, sub.LatestCheckBlockNumber)
	if err != nil {
		return
	}
	newLatestCheckBlock := blocks[0].Number().Int64()
	for _, block := range blocks {
		txs, err := wk.ETH.FilterTransactionFromBlock(block, func(tx *types.Transaction) bool {
			from, _ := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
			isToAddress := false
			if tx.To() != nil {
				isToAddress = tx.To().Hex() == sub.Address
			}
			isFromAddress := from.Hex() == sub.Address
			isSubBeforeTransaction := sub.CreatedAt.Before(tx.Time().UTC())
			return isSubBeforeTransaction && (isToAddress || isFromAddress)
		})
		if err != nil {
			return
		}
		for _, tx := range txs {
			logrus.Infof("Address %s found a transaction %s", sub.Address, tx.Hash().Hex())
			from, _ := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
			// Upsert transaction
			_ = wk.TransactionService.UpsertTransaction(ctx, bson.M{
				"address":         sub.Address,
				"transactionHash": tx.Hash().Hex(),
			}, bson.M{
				"address":         sub.Address,
				"transactionHash": tx.Hash().Hex(),
				"from":            from.Hex(),
				"to":              tx.To().Hex(),
				"value":           new(big.Float).Quo(new(big.Float).SetInt(tx.Value()), big.NewFloat(1e18)).String(),
				"transactionTime": tx.Time().UTC(),
			})
		}
	}
	// Update address newLatestCheckBlock
	if err := wk.SubscribeService.UpdateOne(ctx, bson.M{"address": sub.Address}, bson.M{"latestCheckBlockNumber": newLatestCheckBlock, "updatedAt": time.Now().UTC()}); err != nil {
		logrus.Errorf("Update newLatestCheckBlock failed %v", err)
	}
}
