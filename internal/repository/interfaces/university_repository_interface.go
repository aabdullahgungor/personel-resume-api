package interfaces

import "github.com/aabdullahgungor/personal-resume-api/internal/model"

type IUniversityRepository interface {
	GetAllUniversities() ([]model.University, error)
	GetUniversityById(id int) (model.University, error)
	CreateUniversity(university *model.University) error
	EditUniversity(university *model.University) error
	DeleteUniversity(id int) error
}
