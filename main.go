package main

import (
	"github.com/Qmun14/jwtAuth/database"
	"github.com/Qmun14/jwtAuth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
