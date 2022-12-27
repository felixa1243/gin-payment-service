package main

import (
	"github.com/gin-gonic/gin"
	"v1/controllers"
)

func main() {
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	r.Run(":8080")
}
