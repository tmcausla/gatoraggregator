package main
import (
	"fmt"
	"context"
	"time"
	"log"
	"github.com/google/uuid"
	"strings"
	"database/sql"
	"github.com/tmcausla/gatoraggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(cmd.args))
	}

	fmt.Printf("Collecting feeds every %v\n", cmd.args[0])

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error parsing tiume duration: %w", err)
	}

	ticker := time.NewTicker(timeBetweenReqs)

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			log.Printf("error occurred fetching from url: %v\n", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feed from DB: %w", err)
	}

	feed, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v %w", feed.Url, err) 
	}

	for _, item := range rssFeed.Channel.Item {
		// fmt.Printf("- %v\n", item.Title)

		var parsedTime sql.NullTime
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			parsedTime.Time = t
			parsedTime.Valid = true
		}
		
		parsedTitle := sql.NullString{
			String: item.Title,
			Valid: item.Title != "",
		}

		parsedDescription := sql.NullString{
			String: item.Description,
			Valid: item.Description != "",
		}

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: parsedTitle,
			Url: item.Link,
			Description: parsedDescription,
			PublishedAt: parsedTime,
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") { continue }
			log.Printf("error creating new post: %v\n", err)
		}
	}

	return nil
}