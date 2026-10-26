package main
import (
	"fmt"
	"context"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Printf("error clearing database: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("database successfully reset")
	return nil
}