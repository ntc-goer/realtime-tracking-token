package subscribe

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

func (s *Service) Subscribe(ctx context.Context, addr string) error {
	return s.Repository.Subscribe(ctx, addr)
}

func (s *Service) UnSubscribe(ctx context.Context, addr string) error {
	return s.Repository.UnSubscribe(ctx, addr)
}

func (s *Service) GetAll(ctx context.Context) ([]Subscribe, error) {
	return s.Repository.GetAll(ctx)
}

func (s *Service) UpdateOne(ctx context.Context, filter bson.M, update bson.M) error {
	err := s.Repository.UpdateOne(ctx, filter, update)
	return err
}
