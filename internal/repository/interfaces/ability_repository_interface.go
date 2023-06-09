package interfaces

import "github.com/aabdullahgungor/personal-resume-api/internal/model"

type IAbilityRepository interface {
	GetAllAbilities() ([]model.Ability, error)
	GetAbilityById(id int) (model.Ability, error)
	CreateAbility(ability *model.Ability) error
	EditAbility(ability *model.Ability) error
	DeleteAbility(id int) error
}
