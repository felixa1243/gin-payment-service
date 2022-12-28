package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
	"v1/config"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null;unique" json:"username"`
	Password string `gorm:"type:text;not null" json:"-"`
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}
func (user *User) Save() (*User, error) {
	var err error
	err = config.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
