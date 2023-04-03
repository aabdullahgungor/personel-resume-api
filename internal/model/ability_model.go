package model

import "gorm.io/gorm"

type Ability struct {
	gorm.Model
	AbilityName     string `json:"ability" gorm:"column:ability;not null" `
}