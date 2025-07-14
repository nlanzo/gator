package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/nlanzo/gator/internal/config"
	"github.com/nlanzo/gator/internal/database"

	_ "github.com/lib/pq"
)



type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()
	
	dbQueries := database.New(db)

	app_state := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	commands := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerListUsers)
	commands.register("agg", handlerAgg)
	
	if len(os.Args) < 2 {
		log.Fatal("Usage: gator <command> <args...>")

	}
	args := os.Args[1:]
	err = commands.run(app_state, command{Name: args[0], Args: args[1:]})
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}
}