package Controllers

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello, World! This is Fiber!",
	})
}
