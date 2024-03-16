package middleware

import (
	"github.com/Qmun14/jwtAuth/database"
	"github.com/Qmun14/jwtAuth/models"
	"github.com/gofiber/fiber/v2"
)

func IsEmailHasUsed(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var email models.VerifyEmail

	database.DB.Where("email = ?", data["email"]).First(&email)

	if email.IsUsed || email.Email == data["email"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "email has been used by another account",
		})
	}

	return c.Next()
}

func IsEmailVerified(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)
	if !user.IsEmailVerified {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Please verify your email account",
		})
	}

	return c.Next()
}
