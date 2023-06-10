package interfaces

import "github.com/aabdullahgungor/personal-resume-api/internal/model"

type IExperienceService interface {
	GetAll() ([]model.Experience, error)
	GetById(id string) (model.Experience, error)
	Create(experience *model.Experience) error
	Edit(experience *model.Experience) error
	Delete(id string) error
}
