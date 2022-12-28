package models

type RegistInput struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	PasswordVerify string `json:"passwordVerify" binding:"required"`
}
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
