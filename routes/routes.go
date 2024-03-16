package routes

import (
	"github.com/Qmun14/jwtAuth/controllers"
	"github.com/Qmun14/jwtAuth/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", middleware.IsEmailHasUsed, controllers.Register)
	app.Post("/api/login", middleware.IsEmailVerified, controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/verify_email", controllers.VerifyEmail)
}
