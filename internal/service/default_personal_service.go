package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository/interfaces"
)

var (
	ErrPersonalIDIsNotValid       = errors.New("personal id is not valid")
	ErrPersonalUserNameIsNotEmpty = errors.New("personal username cannot be empty")
	ErrPersonalNotFound           = errors.New("personal cannot be found")
	ErrPersonalEmailIsNotEmpty    = errors.New("personal email cannot be empty")
)

type DefaultPersonalService struct {
	personalRepo interfaces.IPersonalRepository
}

func NewDefaultPersonalService(pRepo interfaces.IPersonalRepository) *DefaultPersonalService {
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
func (d *DefaultPersonalService) GetByEmail(email string) (model.Personal, error) {

	personal, err := d.personalRepo.GetPersonalByEmail(email)

	if err != nil {
		fmt.Println("Service - Personal Not Found")
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
