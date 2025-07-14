package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, function func(*state, command) error) {
	c.registeredCommands[name] = function
}

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command %s not found", cmd.Name)
	}
	return function(s, cmd)
}