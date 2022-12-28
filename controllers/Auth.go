package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"v1/entities"
	"v1/models"
	"v1/utils"
)

func Register(c *gin.Context) {
	var input models.RegistInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if input.Password != input.PasswordVerify {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password and verify password not match",
		})
		return
	}
	user := entities.User{}
	user.Username = input.Username
	user.Password = input.Password
	_, err := user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "registration success",
	})
}
func Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	user := entities.User{}
	user.Username = input.Username
	user.Password = input.Password
	token, err := entities.LoginCheck(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password is incorrect",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
func CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := entities.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
