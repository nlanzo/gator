package main

import (
	"fmt"

	"github.com/nlanzo/gator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("error reading config: %v", err)
	}

	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}

	err = cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	s.cfg = &cfg
	fmt.Printf("Logged in as %s\n", cfg.CurrentUserName)
	
	return nil
}