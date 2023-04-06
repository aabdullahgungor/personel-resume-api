package service

import (
	"errors"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
)

type IPersonalService interface {
	GetAll() ([]model.Personal, error)
	GetById(id string) (model.Personal, error)
	Create(personal *model.Personal) error
	Edit(personal *model.Personal) error
	Delete(id string) error
}

var (
	ErrPersonalIDIsNotValid       = errors.New("Personal id is not valid")
	ErrPersonalUserNameIsNotEmpty = errors.New("Personal username cannot be empty")
	ErrPersonalNotFound           = errors.New("Personal cannot be found")
	ErrPersonalEmailIsNotEmpty = errors.New("Personal email cannot be empty")
)

type DefaultPersonalService struct {
	personalRepo repository.IPersonalRepository
}

func NewDefaultPersonalService(pRepo repository.IPersonalRepository) *DefaultPersonalService {
	return &DefaultPersonalService{
		personalRepo: pRepo,
	}
}

func (d *DefaultPersonalService) GetAll() ([]model.Personal, error) {
	return d.personalRepo.GetAllPersonals()
}
func (d *DefaultPersonalService) GetById(id string) (model.Personal, error) {

	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return model.Personal{}, ErrPersonalIDIsNotValid
	}
	personal, err := d.personalRepo.GetPersonalById(int_id)

	if err != nil {
		return model.Personal{}, err
	}

	return personal, nil
}
func (d *DefaultPersonalService) Create(personal *model.Personal) error {
	return d.personalRepo.CreatePersonal(personal)
}
func (d *DefaultPersonalService) Edit(personal *model.Personal) error {
	return d.personalRepo.EditPersonal(personal)
}
func (d *DefaultPersonalService) Delete(id string) error {
	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return ErrPersonalIDIsNotValid
	}
	err := d.personalRepo.DeletePersonal(int_id)

	if err != nil {
		return ErrPersonalNotFound
	}

	return nil
}

