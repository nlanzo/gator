package main

import (
	"log"
	"os"

	"github.com/nlanzo/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	app_state := &state{
		cfg: &cfg,
	}

	commands := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)


	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> <args...>")

	}
	args := os.Args[1:]
	err = commands.run(app_state, command{Name: args[0], Args: args[1:]})
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}
}