package scraping

import (
	"context"
	"database/sql"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/BrandonRafaelLovelyno/go-rss/internal/database"
	"github.com/google/uuid"
)

func getFeedXml(feedUrl string) (*[]byte, error) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(feedUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	xml, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &xml, nil
}

func getFeedsFromXml(Xml *[]byte) (*ParsedFeed, error) {
	var parsedFeed ParsedFeed

	err := xml.Unmarshal(*Xml, &parsedFeed)
	if err != nil {
		return nil, err
	}

	log.Printf("got %v items", len(parsedFeed.Channel.Items))

	return &parsedFeed, nil
}

func updatePostedCount(mutex *sync.Mutex, postedCount *int, count *int) {
	mutex.Lock()
	*postedCount += *count
	mutex.Unlock()
}

func postFeedItem(feed *database.Feed, parsedFeed *ParsedFeed, query *database.Queries, postedCount *int, mutex *sync.Mutex) error {
	count := 0
	defer updatePostedCount(mutex, postedCount, &count)

	for _, item := range parsedFeed.Channel.Items {
		description := getNullableDescription(item.Description)
		pubDate := parsePubDate(item.PubDate)

		_, err := query.CreatePost(context.Background(), database.CreatePostParams{
			FeedID:      feed.ID,
			Title:       item.Title,
			Description: description,
			Url:         item.Link,
			PublishedAt: pubDate,
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})

		count++

		if err != nil {
			return err
		}
	}

	return nil
}

func parsePubDate(pubDate string) sql.NullTime {
	publishedAt := sql.NullTime{}

	if t, err := time.Parse(time.RFC1123Z, pubDate); err == nil {
		publishedAt = sql.NullTime{
			Time:  t,
			Valid: true,
		}
	}

	return publishedAt
}

func getNullableDescription(description string) sql.NullString {
	if description == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		String: description,
		Valid:  true,
	}
}

func markFeedAsFetched(feed *database.Feed, query *database.Queries) error {
	_, err := query.UpdateFeedLastFetched(context.Background(), database.UpdateFeedLastFetchedParams{
		ID:          feed.ID,
		LastFetched: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}

	return nil
}
