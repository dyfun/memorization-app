package Controllers

import "github.com/gofiber/fiber/v2"

// @Summary Hello
// @Description Hello, World!
// @Tags Hello
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello, World!"
// @Router /api/index [get]
func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello, World! This is Fiber!aa",
	})
}
