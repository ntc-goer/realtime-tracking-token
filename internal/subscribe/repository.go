package subscribe

import (
	"context"
	"github.com/ntc-goer/parser-exercise/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Repository struct {
	Collection database.CollectionInterface
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Collection: db.Collection(SUBSCRIBE_COLLECTION_NAME),
	}
}

func (r *Repository) GetOne(ctx context.Context, filter bson.M) (*Subscribe, error) {
	var sub Subscribe
	res := r.Collection.FindOne(ctx, filter)
	if err := res.Decode(&sub); err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *Repository) Subscribe(ctx context.Context, addr string) error {
	opts := options.Update().SetUpsert(true)
	_, err := r.Collection.UpdateOne(ctx, bson.M{
		"address": addr,
	}, bson.M{
		"$set": bson.M{
			"address":   addr,
			"createdAt": time.Now().UTC(),
			"updatedAt": time.Now().UTC(),
			"deletedAt": nil,
		},
	}, opts)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UnSubscribe(ctx context.Context, addr string) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{
		"address": addr,
	}, bson.M{
		"$set": bson.M{
			"deletedAt": time.Now().UTC(),
		},
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Subscribe, error) {
	cur, err := r.Collection.Find(ctx, bson.M{"deletedAt": nil})
	if err != nil {
		return nil, err
	}
	subscribers := make([]Subscribe, 0)
	if err := cur.All(ctx, &subscribers); err != nil {
		return nil, err
	}
	return subscribers, nil
}

func (r *Repository) UpdateOne(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := r.Collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	return err
}
