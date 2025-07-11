package main

import (
	"fmt"
	"log"

	"github.com/nlanzo/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	err = cfg.SetUser("nlanzo")
	if err != nil {
		log.Fatalf("Error setting user: %v", 	err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config after set user: %v", err)
	}
	fmt.Printf("Read config after set user: %+v\n", cfg)
}