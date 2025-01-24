package api

import (
	"github.com/BrandonRafaelLovelyno/go-rss/internal/auth"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/feeds"
	"github.com/go-chi/chi/v5"
)

func applyFeedsRouter(router *chi.Mux, query *database.Queries) {
	feedHandler := feeds.NewFeedHandler(query)

	router.Post("/feed", auth.Authenticate(feedHandler.HandleCreateFeed, *query))
	router.Get("/feed/user", auth.Authenticate(feedHandler.HandleGetUserFeeds, *query))
}
