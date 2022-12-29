package models

type CustomerInput struct {
	CustomerName string `json:"customer_name" binding:"required"`
}
