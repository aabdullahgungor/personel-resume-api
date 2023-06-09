package interfaces

import "github.com/aabdullahgungor/personal-resume-api/internal/model"

type IExperienceRepository interface {
	GetAllExperiences() ([]model.Experience, error)
	GetExperienceById(id int) (model.Experience, error)
	CreateExperience(experience *model.Experience) error
	EditExperience(experience *model.Experience) error
	DeleteExperience(id int) error
}
