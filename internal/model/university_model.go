package model

import "gorm.io/gorm"

type University struct {
	gorm.Model
	UniversityName     string `json:"university" gorm:"column:university;not null" `
}