package main

import (
	"time"

	"github.com/BrandonRafaelLovelyno/go-rss/api"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/config"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/scraping"
	_ "github.com/lib/pq"
)

func main() {
	conf := config.Load()
	queries := database.Connect(conf.DBUrl)
	go scraping.StartScraping(100*time.Minute, queries, 3)

	api.ListendAndServe(conf.Port, queries)
}
