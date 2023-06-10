package interfaces

import "github.com/aabdullahgungor/personal-resume-api/internal/model"

type IUniversityService interface {
	GetAll() ([]model.University, error)
	GetById(id string) (model.University, error)
	Create(university *model.University) error
	Edit(university *model.University) error
	Delete(id string) error
}
