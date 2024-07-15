package transaction

import (
	"context"
	"github.com/ntc-goer/parser-exercise/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Collection database.CollectionInterface
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Collection: db.Collection(TRANSACTION_COLLECTION_NAME),
	}
}

func (r *Repository) GetByAddress(ctx context.Context, addr string) ([]*Transaction, error) {
	opts := options.Find().SetSort(bson.D{{"transactionTime", -1}})
	cur, err := r.Collection.Find(ctx, bson.M{"address": addr}, opts)
	if err != nil {
		return nil, err
	}
	transactions := make([]*Transaction, 0)
	if err := cur.All(ctx, &transactions); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *Repository) UpsertTransaction(ctx context.Context, filter bson.M, update bson.M) error {
	opts := options.Update().SetUpsert(true)
	_, err := r.Collection.UpdateOne(ctx, filter, bson.M{"$set": update}, opts)
	return err
}
