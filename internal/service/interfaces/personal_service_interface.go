package interfaces

import "github.com/aabdullahgungor/personal-resume-api/internal/model"

type IPersonalService interface {
	GetAll() ([]model.Personal, error)
	GetById(id string) (model.Personal, error)
	GetByEmail(email string) (model.Personal, error)
	Create(personal *model.Personal) error
	Edit(personal *model.Personal) error
	Delete(id string) error
}
