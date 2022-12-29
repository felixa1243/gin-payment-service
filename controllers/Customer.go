package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"v1/entities"
	"v1/models"
	"v1/utils"
)

func Registration(c *gin.Context) {
	var input models.CustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(input.CustomerName)
	cus := entities.Customer{}
	uid, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cus.CustomerName = input.CustomerName
	user, err := entities.GetUserByID(uid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	cus.User = user
	fmt.Println(cus)
	customer, err := cus.Save()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  "OK",
		"content": customer,
	})
}
