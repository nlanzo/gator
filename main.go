package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlanzo/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	app_state := &state{
		cfg: &cfg,
	}

	commands := &commands{
		name_to_function: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: gator <command> <args...>")
		os.Exit(1)
	}
	err = commands.run(app_state, command{name: args[0], args: args[1:]})
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}
}