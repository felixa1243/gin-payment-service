package controllers

import (
	"github.com/gin-gonic/gin"
	"v1/middleware"
)

func Routes() *gin.Engine {
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", Register)
	public.POST("/login", Login)
	protected := r.Group("/api/user")
	protected.Use(middleware.JwtMiddleware())
	protected.GET("/", CurrentUser)
	return r
}
