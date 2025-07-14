package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.Name)
	}

	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Printf("Logged in as %s\n", s.cfg.CurrentUserName)
	
	return nil
}