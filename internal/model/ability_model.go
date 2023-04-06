package model

type Ability struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	AbilityName string     `gorm:"column:ability;not null" json:"ability"  `
	Personals   []Personal `gorm:"many2many:personal_ability"`
}