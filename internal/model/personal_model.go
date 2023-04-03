package model

import "gorm.io/gorm"

type Personal struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null" `
	Surname  string `json:"surname" gorm:"not null"`
	UserName string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	WorkExperiences []WorkExperience 
	UserType string `json:"usertype" gorm:"not null"`
	Abilities []Ability `gorm:"many2many:personal-ability;"`
	Universities []University `gorm:"many2many:personal-university;"`
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