package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=root password=secret dbname=auth_jwt port=5431 sslmode=disable TimeZone=Asia/Jakarta"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to database!")
	}

	DB = conn
}
