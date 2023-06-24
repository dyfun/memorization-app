package Controllers

import (
	"github.com/dyfun/memorization-app/app/Models"
	"github.com/dyfun/memorization-app/config"
	"github.com/gofiber/fiber/v2"
)

func AddWord(c *fiber.Ctx) error {
	word := new(Models.Word)
	if err := c.BodyParser(word); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result := config.Db.Create(&word)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(word)
}
