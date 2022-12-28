package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"v1/config"
	"v1/controllers"
	"v1/entities"
)

func main() {
	loadEnv()
	loadDb()
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	r.Run(":8080")
}
func loadDb() {
	config.Connect()
	config.Database.AutoMigrate(&entities.User{})
}
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
