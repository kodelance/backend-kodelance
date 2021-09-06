package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Fullname string `gorm:"not null"`
}

type UserInput struct {
	Email    string `binding:"required, email"`
	Password string `binding:"required"`
	Fullname string `binding:"required"`
}

type UserOutput struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Token    string `json:"token"`
}

type LoginInput struct {
	Email    string `binding:"required, email"`
	Password string `binding:"required"`
}
