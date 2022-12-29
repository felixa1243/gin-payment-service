package container

import (
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"v1/config"
	"v1/controllers"
	"v1/entities"
)

func Run() {
	loadEnv()
	connectDb()
	r := controllers.Routes()
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/docs.json"),
		ginSwagger.DefaultModelsExpandDepth(-1),
	)
	r.Run(":8080")
}
func connectDb() {
	config.Connect()
	config.Database.AutoMigrate(&entities.Customer{}, &entities.Merchant{}, &entities.User{}, &entities.Wallet{})
}
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
