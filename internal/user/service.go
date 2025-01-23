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

func (s *UserService) Create(ctx context.Context, params database.CreateUserParams) (*database.User, error) {
	user, err := s.Query.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
