package service

import (
	"errors"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository/interfaces"
)

var (
	ErrAbilityIDIsNotValid = errors.New("ability id is not valid")
	ErrAbilityNotFound     = errors.New("ability cannot be found")
)

type DefaultAbilityService struct {
	abilityRepo interfaces.IAbilityRepository
}

func NewDefaultAbilityService(aRepo interfaces.IAbilityRepository) *DefaultAbilityService {
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
		return model.Ability{}, ErrAbilityIDIsNotValid
	}
	ability, err := d.abilityRepo.GetAbilityById(int_id)

	if err != nil {
		return model.Ability{}, err
	}

	return ability, nil
}
func (d *DefaultAbilityService) Create(ability *model.Ability) error {
	return d.abilityRepo.CreateAbility(ability)
}
func (d *DefaultAbilityService) Edit(ability *model.Ability) error {
	return d.abilityRepo.EditAbility(ability)
}
func (d *DefaultAbilityService) Delete(id string) error {
	int_id, errId := strconv.Atoi(id)
	if errId != nil {
		return ErrAbilityIDIsNotValid
	}
	err := d.abilityRepo.DeleteAbility(int_id)

	if err != nil {
		return ErrAbilityNotFound
	}

	return nil
}
