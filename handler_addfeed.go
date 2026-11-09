package main
import (
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/tmcausla/gatoraggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("expected 2 arguments, got %d", len(cmd.args))
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: feedName,
		Url: feedURL,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating new feed: %w", err)
	}

	fmt.Printf("new feed record:\nid: %v\ncreated at: %v\nupdated at: %v\nname: %v\nurl: %v\nuser id: %v",
		feed.ID, feed.CreatedAt, feed.UpdatedAt, feed.Name, feed.Url, feed.UserID)

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	return nil
}