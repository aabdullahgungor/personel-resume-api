package service

import (
	"errors"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository/interfaces"
)

var (
	ErrUniversityIDIsNotValid = errors.New("university id is not valid")
	ErrUniversityNotFound     = errors.New("university cannot be found")
)

type DefaultUniversityService struct {
	universityRepo interfaces.IUniversityRepository
}

func NewDefaultUniversityService(uRepo interfaces.IUniversityRepository) *DefaultUniversityService {
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
