package service

import (
	"errors"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
)

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
		return model.WorkExperience{}, errId
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
	err := d.workExperienceRepo.EditWorkExperience(workExperience)
	if err != nil {
		return ErrWorkExperienceNotFound
	}

	return nil
}
func (d *DefaultWorkExperienceService) Delete(id string) error {
	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return errId
	}
	err := d.workExperienceRepo.DeleteWorkExperience(int_id)

	if err != nil {
		return ErrWorkExperienceNotFound
	}

	return nil
}

