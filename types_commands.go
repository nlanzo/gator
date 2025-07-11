package main

import (
	"fmt"

	"github.com/nlanzo/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	name_to_function map[string]func(*state, command) error
}

func (c *commands) register(name string, function func(*state, command) error) {
	c.name_to_function[name] = function
}

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.name_to_function[cmd.name]
	if !ok {
		return fmt.Errorf("command %s not found", cmd.name)
	}
	return function(s, cmd)
}