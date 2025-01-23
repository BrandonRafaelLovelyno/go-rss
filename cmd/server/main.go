package main

import (
	"github.com/BrandonRafaelLovelyno/go-rss/api"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/config"
	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	conf := config.Load()
	queries := database.Connect(conf.DBUrl)
	api.ListendAndServe(conf.Port, queries)
}
