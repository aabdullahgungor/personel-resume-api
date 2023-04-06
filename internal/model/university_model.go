package model

type University struct {
	ID             uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UniversityName string     `gorm:"column:university;not null" json:"university"`
	Personals      []Personal `gorm:"many2many:personal_university"`
}