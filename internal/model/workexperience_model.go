package model

import "gorm.io/gorm"

type WorkExperience struct {
	gorm.Model
	CompanyName string  `json:"company-name" gorm:"not null" `
	Position string `json:"position"`
	StartYear string `json:"startyear"`
	FinishYear string `json:"finishyear"`
}