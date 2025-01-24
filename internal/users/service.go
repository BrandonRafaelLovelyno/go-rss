package users

import (
	"context"
	"time"

	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/google/uuid"
)

type UserService struct {
	Query *database.Queries
}

func newUserService(query *database.Queries) *UserService {
	return &UserService{
		Query: query,
	}
}

func (s *UserService) create(ctx context.Context, name string) (*database.User, error) {
	user, err := s.Query.CreateUser(ctx, database.CreateUserParams{
		Name:      name,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
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
