package main

import (
	"context"
	"database/sql"
	"strings"
	"fmt"
	"time"
	"log"

	"github.com/google/uuid"
	"github.com/ppllama/gator/internal/database"
	"github.com/ppllama/gator/internal/rss_feed"
)

func handlerAgg(s *state, cmd command) error {
	
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: gator agg <interval time>")
	}

	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error parsing time: %v", err)
	}

	fmt.Printf("Collecting feeds every %s\n", duration.String())

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
    	scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) error {
	dbFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching next feed %v", err)
	}

	markFetchedParams := database.MarkFeedFetchedParams{
		ID: dbFeed.ID,
		LastFetchedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
	}

	if err := s.db.MarkFeedFetched(context.Background(), markFetchedParams); err != nil {
		return fmt.Errorf("error setting null: %v", err)
	}

	feed, err := rssfeed.FetchFeed(context.Background(), dbFeed.Url)
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	for _, item := range(feed.Channel.Item) {

		pubDate := sql.NullTime{
			Time: time.Time{},
			Valid: false,
		}
		tempPubDate, err := parsePubDate(item.PubDate)
		if err == nil {
			pubDate = sql.NullTime{
				Time: tempPubDate,
				Valid: true,
			}	
		}

		postParams := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{
				String: item.Title,
				Valid: true,
			},
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid: true,
			},
			PublishedAt: pubDate,
			FeedID: dbFeed.ID,
		}
		_, err = s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
        		continue // duplicate URL -> skip
			}
			log.Printf("error creating post: %v", err)
		}

		fmt.Printf("New Post Added: %s\n", item.Title)
		fmt.Println()
	}

	return nil
}

// func printFeed(feed *rssfeed.RSSFeed) {

// 	fmt.Printf("Feed Name: %s\n", feed.Channel.Title)
// 	fmt.Printf("Feed Description: %s\n", feed.Channel.Description)

// 	for i, item := range(feed.Channel.Item) {
// 		fmt.Printf("Article %d: %s\n", i + 1, item.Title)
// 		fmt.Printf("Description: %s\n", item.Description)
// 		fmt.Printf("Link: %s\n", item.Link)
// 		fmt.Printf("Date: %s\n", item.PubDate)
// 	}
	
// }

func parsePubDate(s string) (time.Time, error) {
    layouts := []string{
        time.RFC1123Z,
        time.RFC1123,
        time.RFC3339,
        time.RFC3339Nano,
        "Mon, 02 Jan 2006 15:04:05 -0700",
        "Mon, 02 Jan 2006 15:04:05 MST",
    }

    var lastErr error
    for _, layout := range layouts {
        t, err := time.Parse(layout, s)
        if err == nil {
            return t, nil
        }
        lastErr = err
    }
    return time.Time{}, lastErr
}