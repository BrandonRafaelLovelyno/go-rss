package api

import (
	"github.com/BrandonRafaelLovelyno/go-rss/internal/user"
	"github.com/BrandonRafaelLovelyno/go-rss/pkg/utils"
	"github.com/go-chi/chi/v5"
)

func applyUserRoutes(router *chi.Mux, userHandler *user.UserHandler) {
	router.Post("/user", utils.SendResponse(userHandler.HandleCreateUser))
}
