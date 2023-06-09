package repository

import (
	"errors"
	"log"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/database"
	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"gorm.io/gorm"
)

var (
	ErrAbilityNotFound = errors.New("FromRepository - ability not found")
)

type PostgreSqlAbilityRepository struct {
	connectionPool *gorm.DB
}

func NewPostgreSqlAbilityRepository() *PostgreSqlAbilityRepository {

	db := database.GetDatabase()

	return &PostgreSqlAbilityRepository{
		connectionPool: db,
	}
}

func (p *PostgreSqlAbilityRepository) GetAllAbilities() ([]model.Ability, error) {

	var abilities []model.Ability
	result := p.connectionPool.Find(&abilities)
	if result.Error != nil {
		return []model.Ability{}, ErrAbilityNotFound
	}

	return abilities, nil
}
func (p *PostgreSqlAbilityRepository) GetAbilityById(id int) (model.Ability, error) {

	var ability model.Ability
	result := p.connectionPool.First(&ability, id)
	if result.Error != nil {
		return model.Ability{}, ErrAbilityNotFound
	}

	return ability, nil
}
func (p *PostgreSqlAbilityRepository) CreateAbility(ability *model.Ability) error {

	err := p.connectionPool.Create(&ability).Error

	if err != nil {
		panic(err)
	}

	log.Printf("\ndisplay the ids of the newly inserted objects: %v", ability.ID)

	return err
}
func (p *PostgreSqlAbilityRepository) EditAbility(ability *model.Ability) error {

	err := p.connectionPool.Save(&ability).Error

	if err != nil {
		panic(err)
	}

	log.Printf("\ndisplay the ids of the edited objects: %v", ability.ID)

	return err
}
func (p *PostgreSqlAbilityRepository) DeleteAbility(id int) error {

	err := p.connectionPool.Delete(&model.Ability{}, id).Error

	if err != nil {
		panic(err)
	}

	log.Println("deleting the first result from the search filter\n" + "The id of the deleted document:" + strconv.Itoa(id))

	return err
}
