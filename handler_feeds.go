package main
import (
	"fmt"
	"context"
)

func handlerFeeds(s *state, cmd command) error {
	feedRows, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}

	for _, row := range feedRows {
		fmt.Printf("feed name: %v\nfeed url: %v\nauthor: %v\n\n",
			row.Name, row.Url, row.Username)
	}

	return nil
}