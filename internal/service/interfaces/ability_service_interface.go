package interfaces

import "github.com/aabdullahgungor/personal-resume-api/internal/model"

type IAbilityService interface {
	GetAll() ([]model.Ability, error)
	GetById(id string) (model.Ability, error)
	Create(ability *model.Ability) error
	Edit(ability *model.Ability) error
	Delete(id string) error
}
