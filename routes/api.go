package routes

import (
	"github.com/dyfun/memorization-app/app/Controllers"
	"github.com/dyfun/memorization-app/app/Middleware"
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

	// Word routes
	word := api.Group("/word").Use(Middleware.Auth)
	word.Post("/add", Controllers.AddWord)

}
