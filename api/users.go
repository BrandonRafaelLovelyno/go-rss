package api

import (
	"github.com/BrandonRafaelLovelyno/go-rss/internal/auth"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/users"
	"github.com/go-chi/chi/v5"
)

func applyUsersRouter(router *chi.Mux, query *database.Queries) {
	user := users.NewUserHandler(query)
	auth := auth.NewAuthMiddleware(query)

	router.Get("/user", auth.Authenticate(user.HandleGetUser))
	router.Post("/user", user.HandleCreateUser)
}
