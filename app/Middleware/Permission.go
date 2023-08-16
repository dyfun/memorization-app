package Middleware

import "github.com/gofiber/fiber/v2"

func Permission(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if len(roles) == 0 {
			return c.Next()
		}

		return c.Next()
	}
}
