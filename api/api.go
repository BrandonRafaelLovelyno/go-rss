package api

import (
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

func ListendAndServe(port string, query *database.Queries) {
	router := chi.NewRouter()

	applyCORS(router)
	applyAllRoutes(router, query)

	startServer(router, port)
}

func applyAllRoutes(router *chi.Mux, query *database.Queries) {
	applyUsersRouter(router, query)
	applyFeedsRouter(router, query)
	applyFollowsFeedsRouter(router, query)
}

func applyCORS(router *chi.Mux) {
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func startServer(router *chi.Mux, port string) {
	log.Print("Starting server on port " + port)

	server := &http.Server{Addr: ":" + port, Handler: router}
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
