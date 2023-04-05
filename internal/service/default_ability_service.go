package service

import (
	"errors"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
)

type IAbilityService interface {
	GetAll() ([]model.Ability, error)
	GetById(id string) (model.Ability, error)
	Create(ability *model.Ability) error
	Edit(ability *model.Ability) error
	Delete(id string) error
}

var ( 
	ErrAbilityIDIsNotValid       = errors.New("Ability id is not valid")
	ErrAbilityNotFound           = errors.New("Ability cannot be found")
)


type DefaultAbilityService struct {
	abilityRepo repository.IAbilityRepository
}

func NewDefaultAbilityService(aRepo repository.IAbilityRepository) *DefaultAbilityService {
	return &DefaultAbilityService{
		abilityRepo: aRepo,
	}
}

func (d *DefaultAbilityService) GetAll() ([]model.Ability, error) {
	return d.abilityRepo.GetAllAbilities()
}
func (d *DefaultAbilityService) GetById(id string) (model.Ability, error) {

	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return model.Ability{}, errId
	}
	ability, err := d.abilityRepo.GetAbilityById(int_id)

	if err != nil {
		return model.Ability{}, ErrAbilityNotFound
	}

	return ability, nil
}
func (d *DefaultAbilityService) Create(ability *model.Ability) error {
	return d.abilityRepo.CreateAbility(ability)
}
func (d *DefaultAbilityService) Edit(ability *model.Ability) error {
	err := d.abilityRepo.EditAbility(ability)
	if err != nil {
		return ErrAbilityNotFound
	}

	return nil
}
func (d *DefaultAbilityService) Delete(id string) error {
	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return errId
	}
	err := d.abilityRepo.DeleteAbility(int_id)

	if err != nil {
		return ErrAbilityNotFound
	}

	return nil
}

