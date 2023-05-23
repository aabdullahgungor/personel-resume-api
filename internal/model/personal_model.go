package model

import "golang.org/x/crypto/bcrypt"

type Personal struct {
	ID           uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string       `gorm:"column:name;not null" json:"name" `
	Surname      string       `gorm:"column:surname;not null" json:"surname" `
	UserName     string       `gorm:"column:username;not null;unique" json:"username"`
	Email        string       `gorm:"column:email;not null;unique" json:"email"`
	Password     string       `gorm:"column:password;not null" json:"password"`
	Experiences  []Experience `gorm:"foreignKey:PersonalID;references:ID" json:"work_experiences"`
	UserType     string       `gorm:"column:usertype;not null" json:"usertype"`
	Abilities    []Ability    `gorm:"many2many:personal_ability"`
	Universities []University `gorm:"many2many:personal_university"`
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

func (personal *Personal) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	personal.Password = string(bytes)
	return nil
}
func (personal *Personal) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(personal.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
