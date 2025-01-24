package feeds

import (
	"context"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
)

type FeedService struct {
	Query *database.Queries
}

func newFeedService(query *database.Queries) *FeedService {
	return &FeedService{
		Query: query,
	}
}

func (s *FeedService) create(ctx context.Context, params database.CreateFeedParams) (*database.Feed, error) {
	user, err := s.Query.CreateFeed(ctx, params)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *FeedService) getByUser(ctx context.Context, user database.User) ([]database.Feed, error) {
	feeds, err := s.Query.GetFeedByUserId(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}

func (s *FeedService) getFollowedFeedByUser(ctx context.Context, user database.User) ([]database.Feed, error) {
	feeds, err := s.Query.GetFollowedFeedsByUserId(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}
