package Controllers

import (
	"github.com/dyfun/memorization-app/app/Helper"
	"github.com/dyfun/memorization-app/app/Models"
	"github.com/dyfun/memorization-app/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func UserCreate(c *fiber.Ctx) error {
	user := new(Models.UserRegister)
	newUser := new(Models.User)

	// Set body to user struct
	if err := c.BodyParser(user); err != nil {
		return err
	}

	// Check if user exists
	if err := config.Db.Where("email = ?", user.Email).First(&newUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	// Create user
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.Email = user.Email
	newUser.Password = Helper.EncryptPassword(user.Password)

	config.Db.Create(&newUser)

	// Send welcome mail
	to := []string{"gulertayfun@outlook.com"}
	var mailData = struct {
		Name string
	}{
		Name: "Tayfun",
	}
	template := "resources/views/mail/register.html"
	Helper.SendEmail(to, "Welcome!", template, mailData)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created",
	})
}

func UserLogin(c *fiber.Ctx) error {
	// Body parser
	request := new(Models.UserLogin)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Find user in database and assign to existUser
	var user Models.User
	if err := config.Db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "E-mail or password wrong",
		})
	}

	// Compare password
	checkPassword := Helper.DecryptPassword(request.Password, user.Password)
	if checkPassword == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "E-mail or password wrong",
		})
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
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
