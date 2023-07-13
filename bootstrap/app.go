package bootstrap

import (
	"fmt"
	"github.com/dyfun/memorization-app/routes"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	fmt.Println("Project started")

	// Create a new Fiber instance
	app := fiber.New()

	// Register all routes
	routes.AllRoutes(app)

	// Start server
	app.Listen(":8000")
}
