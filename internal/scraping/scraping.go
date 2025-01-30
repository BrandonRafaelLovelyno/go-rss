package scraping

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
)

func StartScraping(interval time.Duration, query *database.Queries, thread int) {
	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
		log.Printf("starting scraping %v at %v", thread, time.Now())

		wg := &sync.WaitGroup{}

		feedsToFetch, err := query.GetFeedsToFetch(context.Background(), int32(thread))
		if err != nil {
			log.Fatal("failed to get feeds to fetch: ", err)
		}

		mutex := &sync.Mutex{}
		postedCount := 0
		for _, feed := range feedsToFetch {
			wg.Add(1)
			go scrapeFeed(wg, query, &feed, &postedCount, mutex)
		}

		wg.Wait()

		log.Printf("finished scraping %v feeds with %v new posts at %v", len(feedsToFetch), postedCount, time.Now())
	}
}

func scrapeFeed(wg *sync.WaitGroup, query *database.Queries, feed *database.Feed, postedCount *int, mutex *sync.Mutex) {
	defer wg.Done()

	xml, err := getFeedXml(feed.Url)
	if err != nil {
		log.Printf("failed to fetch feed xml: %s", err.Error())
		return
	}

	parsedFeed, err := getFeedsFromXml(xml)
	if err != nil {
		log.Printf("failed to parse feed xml: %s", err.Error())
		return
	}

	err = postFeedItem(feed, parsedFeed, query, postedCount, mutex)
	if err != nil && !strings.Contains(err.Error(), "posts_title_key") {
		log.Printf("failed to post feed item: %s", err.Error())
		return
	}

	err = markFeedAsFetched(feed, query)
	if err != nil {
		log.Printf("failed to mark feed as fetched: %s", err.Error())
		return
	}
}

type ParsedFeed struct {
	Channel struct {
		Title       string           `xml:"title"`
		Link        string           `xml:"link"`
		Description string           `xml:"description"`
		Language    string           `xml:"language"`
		Items       []ParsedFeedItem `xml:"item"`
	} `xml:"channel"`
}

type ParsedFeedItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
