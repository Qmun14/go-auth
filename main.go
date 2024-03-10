package main

import (
	"log"

	"github.com/Qmun14/jwtAuth/database"
	"github.com/Qmun14/jwtAuth/routes"
	"github.com/Qmun14/jwtAuth/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("tidak bisa memuat config:", err)
	}

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.ClientAddress,
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(config.ServerAddress)
}
