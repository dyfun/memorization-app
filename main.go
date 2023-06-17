package main

import (
	"github.com/dyfun/memorization-app/config"
	"github.com/dyfun/memorization-app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load .env file
	config.LoadEnv()

	// Connect to database
	config.Connect()

	// Start server
	app := fiber.New()

	// Register all routes
	routes.AllRoutes(app)

	app.Listen(":8000")
}
