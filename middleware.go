package main
import (
	"fmt"
	"context"
	"github.com/tmcausla/gatoraggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		curUser, err := s.db.GetUser(context.Background(), s.cfg.CurUserName)
		if err != nil {
			return fmt.Errorf("error fetching user: %w", err)
		}
		return handler(s, cmd, curUser)
	}
}