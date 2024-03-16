package controllers

import (
	"log"
	"strconv"
	"time"

	"github.com/Qmun14/jwtAuth/database"
	"github.com/Qmun14/jwtAuth/models"
	"github.com/Qmun14/jwtAuth/services"
	"github.com/Qmun14/jwtAuth/utils"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := utils.HashPassword(data["password"])
	secretCode := utils.RandomString(32)

	user := models.User{
		Id:       uuid.New(),
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	verifyEmail := models.VerifyEmail{
		Email:      data["email"],
		SecretCode: secretCode,
	}

	err := database.DB.Create(&user).Error
	if err != nil {
		log.Fatal("data tidak berhasil di simpan karena 1111!", err)
	}

	query := "INSERT INTO verify_emails (email, secret_code) VALUES (?, ?)"
	err = database.DB.Exec(query, data["email"], secretCode).Error
	if err != nil {
		log.Fatal("data tidak berhasil di simpan karena! ", err)
	}

	query = "SELECT id ,email FROM verify_emails WHERE email = ?"
	row := database.DB.Raw(query, data["email"]).Row()
	row.Scan(&verifyEmail.ID, &verifyEmail.Email)

	if err := services.CreateVerifyEmail(verifyEmail.ID, verifyEmail.SecretCode, verifyEmail.Email, user.Name); err != nil {
		log.Fatalf("email tidak berhasil terkirim karena: %v", err)
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	config, err := utils.LoadConfig("..")
	if err != nil {
		log.Fatal("tidak bisa memuat config:", err)
	}
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == uuid.Nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := utils.CheckPassword(data["password"], user.Password); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Id.String(),
		ExpiresAt: jwt.NewTime(float64(time.Now().Add(24 * time.Hour).Unix())),
	})

	token, err := claims.SignedString([]byte(config.SecretCode))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func User(c *fiber.Ctx) error {
	config, err := utils.LoadConfig("..")
	if err != nil {
		log.Fatal("tidak bisa memuat config:", err)
	}
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SecretCode), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func VerifyEmail(c *fiber.Ctx) error {
	emailId, _ := strconv.Atoi(c.Query("email_id"))
	secret_code := c.Query("secret_code")

	var email string
	var is_used bool

	row := database.DB.Raw("select email, is_used from verify_emails where id = ?", emailId).Row()
	row.Scan(&email, &is_used)

	if is_used {
		return c.JSON(fiber.Map{
			"message": "failed to verify email",
		})
	}

	err := database.DB.Table("verify_emails").Where("id = ?", int64(emailId)).Where("secret_code = ?", secret_code).Where("is_used = ?", false).Update("is_used", true).Error
	if err != nil {
		log.Fatal("data tidak berhasil di Update karena! ", err)
		return c.JSON(fiber.Map{
			"message": "failed to verify email",
		})
	}

	err = database.DB.Table("users").Where("email = ?", email).Update("is_email_verified", true).Error
	if err != nil {
		log.Fatal("data tidak berhasil di Update karena! ", err)
		return c.JSON(fiber.Map{
			"message": "failed to verify email",
		})
	}

	return c.JSON(fiber.Map{
		"message": "email has been verified",
	})
}
