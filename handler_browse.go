package main
import (
	"fmt"
	"log"
	"context"
	"strconv"
	"github.com/tmcausla/gatoraggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		log.Println("too many arguments, expected 0 or 1.  defaulting fetch limit to 2")
	}
	fetchLimit := 2
	if len(cmd.args) == 1 {
		if num, err := strconv.Atoi(cmd.args[0]); err == nil {
			fetchLimit = num
		} else {
			log.Println("optional limit parameter must be int, defaulting to 2")
		}
	}
	
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(fetchLimit),
	})
	if err != nil {
		return fmt.Errorf("error fetching posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("%v (%v)\n", post.Title.String, post.PublishedAt.Time)
		fmt.Printf("- %v\n", post.Description.String)
		fmt.Printf("from: %v\n\n", post.Url)
	}

	return nil
}