package controllers

import (
	"github.com/gin-gonic/gin"
	"v1/middleware"
)

func Routes() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	auth.POST("/register", Register)
	auth.POST("/login", Login)

	protected := r.Group("/api/user")
	protected.Use(middleware.JwtMiddleware())
	protected.GET("/", CurrentUser)

	customer := r.Group("/api/customer")
	customer.Use(middleware.JwtMiddleware())
	customer.POST("/register", Registration)
	return r
}
