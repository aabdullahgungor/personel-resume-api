package service

import (
	"errors"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository/interfaces"
)

var (
	ErrExperienceIDIsNotValid = errors.New("experience id is not valid")
	ErrExperienceNotFound     = errors.New("experience cannot be found")
)

type DefaultExperienceService struct {
	experienceRepo interfaces.IExperienceRepository
}

func NewDefaultExperienceService(eRepo interfaces.IExperienceRepository) *DefaultExperienceService {
	return &DefaultExperienceService{
		experienceRepo: eRepo,
	}
}

func (d *DefaultExperienceService) GetAll() ([]model.Experience, error) {
	return d.experienceRepo.GetAllExperiences()
}
func (d *DefaultExperienceService) GetById(id string) (model.Experience, error) {

	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return model.Experience{}, ErrExperienceIDIsNotValid
	}
	experience, err := d.experienceRepo.GetExperienceById(int_id)

	if err != nil {
		return model.Experience{}, ErrExperienceNotFound
	}

	return experience, nil
}
func (d *DefaultExperienceService) Create(experience *model.Experience) error {
	return d.experienceRepo.CreateExperience(experience)
}
func (d *DefaultExperienceService) Edit(experience *model.Experience) error {
	return d.experienceRepo.EditExperience(experience)
}
func (d *DefaultExperienceService) Delete(id string) error {
	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return ErrExperienceIDIsNotValid
	}
	err := d.experienceRepo.DeleteExperience(int_id)

	if err != nil {
		return ErrExperienceNotFound
	}

	return nil
}
