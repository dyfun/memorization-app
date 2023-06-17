package routes

import (
	"github.com/dyfun/memorization-app/app/Controllers"
	"github.com/gofiber/fiber/v2"
)

func AllRoutes(app *fiber.App) {
	app.Get("/", Controllers.Index)
}
