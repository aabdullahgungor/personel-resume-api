package service

import (
	"errors"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
)

type IUniversityService interface {
	GetAll() ([]model.University, error)
	GetById(id string) (model.University, error)
	Create(university *model.University) error
	Edit(university *model.University) error
	Delete(id string) error
}

var ( 
	ErrUniversityIDIsNotValid       = errors.New("University id is not valid")
	ErrUniversityNotFound           = errors.New("University cannot be found")
)


type DefaultUniversityService struct {
	universityRepo repository.IUniversityRepository
}

func NewDefaultUniversityService(uRepo repository.IUniversityRepository) *DefaultUniversityService {
	return &DefaultUniversityService{
		universityRepo: uRepo,
	}
}

func (d *DefaultUniversityService) GetAll() ([]model.University, error) {
	return d.universityRepo.GetAllUniversities()
}
func (d *DefaultUniversityService) GetById(id string) (model.University, error) {

	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return model.University{}, ErrUniversityIDIsNotValid
	}
	university, err := d.universityRepo.GetUniversityById(int_id)

	if err != nil {
		return model.University{}, ErrUniversityNotFound
	}

	return university, nil
}
func (d *DefaultUniversityService) Create(university *model.University) error {
	return d.universityRepo.CreateUniversity(university)
}
func (d *DefaultUniversityService) Edit(university *model.University) error {
	return d.universityRepo.EditUniversity(university)
}
func (d *DefaultUniversityService) Delete(id string) error {
	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return ErrUniversityIDIsNotValid
	}
	err := d.universityRepo.DeleteUniversity(int_id)

	if err != nil {
		return ErrUniversityNotFound
	}

	return nil
}

