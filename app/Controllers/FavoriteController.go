package Controllers

import (
	"github.com/dyfun/memorization-app/app/Models"
	"github.com/dyfun/memorization-app/config"
	"github.com/gofiber/fiber/v2"
)

func AddFavorite(c *fiber.Ctx) error {
	word := new(Models.Word)
	favorite := new(Models.FavoriteAdd)

	if err := c.BodyParser(favorite); err != nil {
		return err
	}

	checkWord := config.Db.Where("id = ?", favorite.WordID).First(&word)
	if checkWord.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Word not found",
		})
	}

	favoriteItem := Models.Favorite{
		UserID: favorite.UserID,
		WordID: favorite.WordID,
	}

	result := config.Db.Create(&favoriteItem)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	err := config.Db.Preload("User").Preload("Word").First(&favoriteItem).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(favoriteItem)
}
