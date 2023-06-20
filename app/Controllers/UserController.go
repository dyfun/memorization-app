package Controllers

import (
	"github.com/dyfun/memorization-app/app/Models"
	"github.com/dyfun/memorization-app/config"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserCreate(c *fiber.Ctx) error {
	user := new(Models.User)

	// Set body to user struct
	if err := c.BodyParser(user); err != nil {
		return err
	}

	// Check if user exists
	if err := config.Db.Where("email = ?", user.GetEmail()).First(&user).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	// Hashing the password
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.SetPassword(string(hash))

	// Create user
	config.Db.Create(&user)

	return nil
}
