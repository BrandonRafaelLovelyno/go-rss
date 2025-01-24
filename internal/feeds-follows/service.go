package feedsfollows

import (
	"context"
	"time"

	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/google/uuid"
)

type FeedFollowService struct {
	query *database.Queries
}

func newFeedFollowService(query *database.Queries) *FeedFollowService {
	return &FeedFollowService{
		query: query,
	}
}

func (s *FeedFollowService) createFeedFollow(ctx context.Context, userID, feedID uuid.UUID) error {
	_, err := s.query.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	})

	return err
}
