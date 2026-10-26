package main
import (
	"fmt"
	"github.com/tmcausla/gatoraggregator/internal/config"
	"github.com/tmcausla/gatoraggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if runCmd, ok := c.cmds[cmd.name]; ok {
		return runCmd(s, cmd)
	}
	return fmt.Errorf("unrecognized command: %s", cmd.name)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}