package transaction

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Service struct {
	Repository *Repository
}

func NewService(r *Repository) *Service {
	return &Service{
		Repository: r,
	}
}

func (s *Service) GetByAddress(ctx context.Context, addr string) ([]*Transaction, error) {
	return s.Repository.GetByAddress(ctx, addr)
}

func (s *Service) UpsertTransaction(ctx context.Context, filter bson.M, update bson.M) error {
	return s.Repository.UpsertTransaction(ctx, filter, update)
}
