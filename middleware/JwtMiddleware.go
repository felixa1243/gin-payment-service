package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"v1/utils"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
