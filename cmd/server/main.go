package main

import (
	"database/sql"
	"github.com/BrandonRafaelLovelyno/go-rss/handler"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type apiCfg struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("PG_URL")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if dbUrl == "" {
		log.Fatal("$PG_URL must be set")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Error making pool connection: ", err)
	}
	if err = conn.Ping(); err != nil {
		log.Fatal("Cannot ping database: ", err)
	}

	//	api := &apiCfg{DB: database.New(conn)}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/", handler.HandleReadiness)

	log.Print("Starting server on port " + port)
	server := &http.Server{Addr: ":" + port, Handler: router}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
