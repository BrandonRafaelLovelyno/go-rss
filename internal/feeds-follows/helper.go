package feedsfollows

import (
	"errors"
	"net/http"

	"github.com/BrandonRafaelLovelyno/go-rss/pkg/utils"
	"github.com/google/uuid"
)

type FollowFeedParams struct {
	FeedID uuid.UUID `json:"feed_id"`
}

func getFollowFeedParams(r *http.Request) (*FollowFeedParams, error) {
	parameter := FollowFeedParams{}
	if err := utils.ReadParams(r, &parameter); err != nil {
		return &FollowFeedParams{}, errors.New("invalid request body")
	}

	return &parameter, nil
}

func makeMessageResponse(message string) interface{} {
	return struct {
		Message string `json:"message"`
	}{Message: message}
}
