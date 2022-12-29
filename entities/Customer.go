package entities

import (
	"gorm.io/gorm"
	"v1/config"
)

type Customer struct {
	gorm.Model
	CustomerName string
	User         User
}

func (Customer) TableName() string {
	return "customer"
}
func (customer *Customer) Save() (*Customer, error) {
	var err error
	err = config.Database.Create(&customer).Error
	if err != nil {
		return &Customer{}, err
	}

	w := Wallet{}
	w.CustomerId = customer.ID
	w.Account = 0

	err = config.Database.Create(&w).Error

	return customer, nil
}
