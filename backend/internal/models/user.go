package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement:true" json:id`
	Username string `gorm:"unique;not null" json:username`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"-"`
	Picture  string `json:"picture"`
}

type LoginForm struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type SignUpForm struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
