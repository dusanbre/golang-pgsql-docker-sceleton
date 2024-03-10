package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, name, port)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dns,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
