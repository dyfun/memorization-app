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

func GetWords(c *fiber.Ctx) error {
	words := []Models.Word{}
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	pagination := (page - 1) * limit

	result := config.Db.Order("ID desc").Offset(pagination).Limit(limit).Find(&words)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": words,
	})
}

func UpdateWord(c *fiber.Ctx) error {
	id := c.Params("id")

	word := new(Models.Word)
	result := config.Db.First(&word, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Word not found",
		})
	}

	if err := c.BodyParser(word); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	config.Db.Save(&word)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Word updated successfully",
	})
}

func DeleteWord(c *fiber.Ctx) error {
	id := c.Params("id")
	var word Models.Word

	result := config.Db.First(&word, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Word not found",
		})
	}

	config.Db.Delete(&word, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Word deleted successfully",
	})
}
