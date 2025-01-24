package feedsfollows

import (
	"fmt"
	"net/http"

	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/pkg/utils"
)

type FeedFollowHandler struct {
	Service *FeedFollowService
}

func NewFeedFollowHandler(query *database.Queries) *FeedFollowHandler {
	return &FeedFollowHandler{
		Service: newFeedFollowService(query),
	}
}

func (h *FeedFollowHandler) HandleFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	parameter, err := getFollowFeedParams(r)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error reading request: %s", err))
		return
	}

	err = h.Service.createFeedFollow(r.Context(), user.ID, parameter.FeedID)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("error following feed: %s", err))
		return
	}

	msg := makeMessageResponse("Followed feed")
	utils.RespondWithJson(w, 201, msg)
}

func (h *FeedFollowHandler) HandleGetFollowedFeed(w http.ResponseWriter, r *http.Request, user database.User) {
}
