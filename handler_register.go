package main
import (
	"fmt"
	"time"
	"github.com/google/uuid"
	"os"
	"context"
	"github.com/tmcausla/gatoraggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(cmd.args))
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
	})
	if err != nil {
		fmt.Printf("error creating user: %v", err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("user has been set to: %s\n", user.Name)
	//fmt.Printf("UserID: %v\ncreated at: %v\nlast update at: %v\nname: %v\n", 
		//user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	return nil
}