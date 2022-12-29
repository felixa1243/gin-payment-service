package entities

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	MerchantName string
	User         User
	Wallet       Wallet
}

func (Merchant) TableName() string {
	return "m_merchant"
}
