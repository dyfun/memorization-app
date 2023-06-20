package Controllers

import (
	"github.com/dyfun/memorization-app/app/Models"
	"github.com/dyfun/memorization-app/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created",
	})
}

func UserLogin(c *fiber.Ctx) error {
	// Body parser
	request := new(Models.UserLogin)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	// Find user in database and assign to existUser
	var user Models.User
	if err := config.Db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "E-mail or password wrong",
		})
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "E-mail or password wrong",
		})
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.GetEmail(),
		"id":    user.GetId(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign token
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}
