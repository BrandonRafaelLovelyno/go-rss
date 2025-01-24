package api

import (
	"github.com/BrandonRafaelLovelyno/go-rss/internal/auth"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/feeds-follows"
	"github.com/go-chi/chi/v5"
)

func applyFollowsFeedsRouter(router *chi.Mux, query *database.Queries) {
	followFeeds := feedsfollows.NewFeedFollowHandler(query)
	auth := auth.NewAuthMiddleware(query)

	router.Post("/feed/follow", auth.Authenticate(followFeeds.HandleFollowFeed))
}
