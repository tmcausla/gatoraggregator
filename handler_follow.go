package main
import (
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/tmcausla/gatoraggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(cmd.args))
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed-follow: %w", err)
	}

	fmt.Printf("new feed follow created\nfeed name: %v\nfollowed by: %v\n",
		feedFollow.FeedName, feedFollow.UserName)

	return nil
}