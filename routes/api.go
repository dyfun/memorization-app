package routes

import (
	"github.com/dyfun/memorization-app/app/Controllers"
	"github.com/dyfun/memorization-app/app/Middleware"
	_ "github.com/dyfun/memorization-app/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func AllRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	api := app.Group("/api")
	api.Get("/index", Controllers.Index)
	api.Get("/example", Controllers.Example)

	// User routes
	user := api.Group("/user")
	user.Post("/create", Controllers.UserCreate)
	user.Post("/login", Controllers.UserLogin)

	// Word routes
	word := api.Group("/word").Use(Middleware.Auth)
	word.Post("/add", Middleware.HasPermission("word->add"), Controllers.AddWord)
	word.Get("/all", Controllers.GetWords)
	word.Put("/update/:id", Controllers.UpdateWord)
	word.Delete("/delete/:id", Controllers.DeleteWord)

	// Favorite routes
	favorite := api.Group("/favorite").Use(Middleware.Auth)
	favorite.Post("/add", Controllers.AddFavorite)
}
