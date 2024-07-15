package transaction

import "time"

const TRANSACTION_COLLECTION_NAME = "transaction"

type Transaction struct {
	Address         string    `json:"address" bson:"address"`
	TransactionHash string    `json:"transactionHash" bson:"transactionHash"`
	From            string    `json:"from" bson:"from"`
	To              string    `json:"to" bson:"to"`
	Value           string    `json:"value" bson:"value"`
	TransactionTime time.Time `json:"transactionTime" bson:"transactionTime"`
}
