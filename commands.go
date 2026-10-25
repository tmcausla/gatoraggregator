package main
import (
	"fmt"
	"github.com/tmcausla/gatoraggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(cmd.args))
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("user has been set to: %s\n", s.cfg.CurUserName)
	return nil
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