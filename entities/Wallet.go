package entities

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Account    uint `json:"account"`
	MerchantId uint
	CustomerId uint
}
