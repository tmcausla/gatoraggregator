package main
import (
	"fmt"
	"context"
	"github.com/tmcausla/gatoraggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}

	fmt.Printf("%v is currently following:\n", user.Name)
	for _, feed := range followedFeeds {
		fmt.Printf("- %v\n", feed.FeedName)
	}

	return nil
}