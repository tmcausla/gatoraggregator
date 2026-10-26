package main

import (
	"fmt"
	"context"
	"os"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("error getting users: %v", err)
		os.Exit(1)
	}

	for _, user := range users{
		output := fmt.Sprintf("* %s", user.Name)
		if user.Name == s.cfg.CurUserName {
			output += " (current)"
		}
		fmt.Println(output)
	}
	return nil
}