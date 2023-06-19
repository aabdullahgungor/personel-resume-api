package model

type Experience struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	CompanyName string `gorm:"column:company;not null" json:"company-name"  `
	Position    string `gorm:"column:position" json:"position"`
	StartYear   string `gorm:"column:start_year" json:"startyear"`
	FinishYear  string `gorm:"column:finish_year" json:"finishyear"`
	PersonalID  int    `gorm:"column:personal_id" json:"personal_id"`
	Personal    Personal
}
