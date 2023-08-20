package Middleware

import (
	"github.com/dyfun/memorization-app/app/Models"
	"github.com/dyfun/memorization-app/config"
	"github.com/gofiber/fiber/v2"
)

func HasPermission(action string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get user
		email := c.Locals("email").(string)

		var user Models.User
		err := config.Db.Preload("Role.RolePermission.Permission").Where("email = ?", email).First(&user).Error
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "User not found!",
			})
		}

		// Check if user has permission
		if user.Role.Name == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "User has no roles!",
			})
		}

		// Check user permission
		for _, permission := range user.Role.RolePermission {
			if permission.Permission.Name == action {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden!",
		})
	}
}
