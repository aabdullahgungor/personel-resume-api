package model

import "gorm.io/gorm"

type Personal struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null" `
	Surname  string `json:"surname" gorm:"not null"`
	UserName string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	UserType string `json:"usertype" gorm:"not null"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}