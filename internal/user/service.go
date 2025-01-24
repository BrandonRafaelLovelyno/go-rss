package user

import (
	"context"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
)

type UserService struct {
	Query *database.Queries
}

func NewUserService(query *database.Queries) *UserService {
	return &UserService{
		Query: query,
	}
}

func (s *UserService) create(ctx context.Context, params database.CreateUserParams) (*database.User, error) {
	user, err := s.Query.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) getByApiKey(ctx context.Context, apiKey string) (*database.User, error) {
	user, err := s.Query.GetUserByApiKey(ctx, apiKey)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetByApiKey(query database.Queries, ctx context.Context, apiKey string) (*database.User, error) {
	user, err := query.GetUserByApiKey(ctx, apiKey)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
