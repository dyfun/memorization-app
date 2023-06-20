package routes

import (
	"github.com/dyfun/memorization-app/app/Controllers"
	"github.com/gofiber/fiber/v2"
)

func AllRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/index", Controllers.Index)
	api.Get("/example", Controllers.Example)

	// User routes
	user := api.Group("/user")
	user.Post("/create", Controllers.UserCreate)
	user.Post("/login", Controllers.UserLogin)
}
