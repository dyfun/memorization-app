package main

import (
	"flag"
	"fmt"
	"github.com/dyfun/memorization-app/bootstrap"
	"github.com/dyfun/memorization-app/config"
	"github.com/dyfun/memorization-app/database/migrations"
	"github.com/dyfun/memorization-app/database/seeders"
)

func main() {
	fmt.Println("Memorization App Started")

	// Register command line arguments
	flag.Parse()
	args := flag.Args()

	// Load .env file
	config.LoadEnv()

	// Connect to database
	config.Connect()

	// Run command
	if len(args) > 0 {
		switch args[0] {
		case "start":
			bootstrap.Start()
		case "seed":
			seeders.Seed()
		case "migrate":
			migrations.Migrate()
		}
	}
}
