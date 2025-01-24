package api

import (
	"github.com/BrandonRafaelLovelyno/go-rss/internal/auth"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/users"
	"github.com/go-chi/chi/v5"
)

func applyUsersRouter(router *chi.Mux, query *database.Queries) {
	userHandler := users.NewUserHandler(query)

	router.Get("/user", auth.Authenticate(userHandler.HandleGetUser, *query))
	router.Post("/user", userHandler.HandleCreateUser)
}
