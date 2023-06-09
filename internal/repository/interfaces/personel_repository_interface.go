package interfaces

import "github.com/aabdullahgungor/personal-resume-api/internal/model"

type IPersonalRepository interface {
	GetAllPersonals() ([]model.Personal, error)
	GetPersonalById(id int) (model.Personal, error)
	GetPersonalByEmail(email string) (model.Personal, error)
	CreatePersonal(personal *model.Personal) error
	EditPersonal(personal *model.Personal) error
	DeletePersonal(id int) error
}
