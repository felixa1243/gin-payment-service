package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	port := os.Getenv("PORT")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dburl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, dbuser, dbpassword, dbname, port)
	Database, err = gorm.Open(postgres.Open(dburl), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database successfully")
	}
}
