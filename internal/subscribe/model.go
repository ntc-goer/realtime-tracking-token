package subscribe

import "time"

const SUBSCRIBE_COLLECTION_NAME = "subscribe"

type Transaction struct {
}

type Subscribe struct {
	Address                string    `json:"address" bson:"address"`
	LatestCheckBlockNumber int64     `json:"-,omitempty" bson:"latestCheckBlockNumber""`
	CreatedAt              time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt" bson:"updatedAt"`
	DeletedAt              time.Time `json:"deletedAt" bson:"deletedAt"`
}
