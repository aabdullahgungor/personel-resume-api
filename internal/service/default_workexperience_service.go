package service

import (
	"errors"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
)

type IWorkExperienceService interface {
	GetAll() ([]model.WorkExperience, error)
	GetById(id string) (model.WorkExperience, error)
	Create(workExperience *model.WorkExperience) error
	Edit(workExperience *model.WorkExperience) error
	Delete(id string) error
}

var ( 
	ErrWorkExperienceIDIsNotValid       = errors.New("WorkExperience id is not valid")
	ErrWorkExperienceNotFound           = errors.New("WorkExperience cannot be found")
)


type DefaultWorkExperienceService struct {
	workExperienceRepo repository.IWorkExperienceRepository
}

func NewDefaultWorkExperienceService(wRepo repository.IWorkExperienceRepository) *DefaultWorkExperienceService {
	return &DefaultWorkExperienceService{
		workExperienceRepo: wRepo,
	}
}

func (d *DefaultWorkExperienceService) GetAll() ([]model.WorkExperience, error) {
	return d.workExperienceRepo.GetAllWorkExperiences()
}
func (d *DefaultWorkExperienceService) GetById(id string) (model.WorkExperience, error) {

	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return model.WorkExperience{}, ErrWorkExperienceIDIsNotValid
	}
	workExperience, err := d.workExperienceRepo.GetWorkExperienceById(int_id)

	if err != nil {
		return model.WorkExperience{}, ErrWorkExperienceNotFound
	}

	return workExperience, nil
}
func (d *DefaultWorkExperienceService) Create(workExperience *model.WorkExperience) error {
	return d.workExperienceRepo.CreateWorkExperience(workExperience)
}
func (d *DefaultWorkExperienceService) Edit(workExperience *model.WorkExperience) error {
	return d.workExperienceRepo.EditWorkExperience(workExperience)
}
func (d *DefaultWorkExperienceService) Delete(id string) error {
	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return ErrWorkExperienceIDIsNotValid
	}
	err := d.workExperienceRepo.DeleteWorkExperience(int_id)

	if err != nil {
		return ErrWorkExperienceNotFound
	}

	return nil
}

