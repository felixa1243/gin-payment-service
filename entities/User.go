package entities

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
	"v1/config"
	"v1/utils"
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
func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	err = config.Database.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func GetUserByID(uid uint) (User, error) {

	var u User

	if err := config.Database.First(&u, uid).Error; err != nil {
		return u, errors.New("User not exists")
	}

	u.PrepareGive()

	return u, nil

}
func (u *User) PrepareGive() {
	u.Password = ""
}
