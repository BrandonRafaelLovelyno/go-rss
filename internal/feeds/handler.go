package feeds

import (
	"fmt"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/pkg/utils"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type FeedHandler struct {
	Service *FeedService
}

func NewFeedHandler(query *database.Queries) *FeedHandler {
	return &FeedHandler{
		Service: newFeedService(query),
	}
}

func (h *FeedHandler) HandleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	parameter, err := getCreateFeedParams(r)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error reading request: %s", err))
	}

	feed, err := h.Service.create(r.Context(), database.CreateFeedParams{
		Name:      parameter.Name,
		Url:       parameter.Url,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("error creating feed: %s", err))
	}

	utils.RespondWithJson(w, 201, feed)
}

func (h *FeedHandler) HandleGetUserFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := h.Service.getByUser(r.Context(), user)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("error getting feeds: %s", err))
	}

	utils.RespondWithJson(w, 200, feeds)
}

func (h *FeedHandler) HandleGetFollowedFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := h.Service.getFollowedFeedByUser(r.Context(), user)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("error getting feed: %s", err))
	}

	utils.RespondWithJson(w, 200, feeds)
}
