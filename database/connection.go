package database

import (
	"fmt"
	"log"

	"github.com/Qmun14/jwtAuth/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config, err := utils.LoadConfig("..")
	if err != nil {
		log.Fatalf("error %s", err)
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to database!")
	}

	DB = conn
}
