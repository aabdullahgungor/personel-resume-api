package repository

import (
	"errors"
	"log"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/database"
	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"gorm.io/gorm"
)

type IExperienceRepository interface {
	GetAllExperiences() ([]model.Experience, error)
	GetExperienceById(id int) (model.Experience, error)
	CreateExperience(experience *model.Experience) error
	EditExperience(experience *model.Experience) error
	DeleteExperience(id int) error
}

var (
	ErrExperienceNotFound = errors.New("FromRepository - experience  not found")
)

type PostgreSqlExperienceRepository struct {
	connectionPool *gorm.DB
}

func NewPostgreSqlExperienceRepository() *PostgreSqlExperienceRepository {

	db := database.GetDatabase()

	return &PostgreSqlExperienceRepository{
		connectionPool: db,
	}
}
func (p *PostgreSqlExperienceRepository) GetAllExperiences() ([]model.Experience, error){

	var experiences []model.Experience
	result := p.connectionPool.Preload("Personal").Find(&experiences)
	if result.Error != nil {
        return []model.Experience{}, ErrExperienceNotFound
	}

	return experiences, nil
}
func (p *PostgreSqlExperienceRepository) GetExperienceById(id int) (model.Experience, error) {

	var experience model.Experience
	result := p.connectionPool.Preload("Personal").First(&experience, id)
	if result.Error != nil {
		return model.Experience{}, ErrExperienceNotFound
	}

	return experience, nil
}
func (p *PostgreSqlExperienceRepository) CreateExperience(experience *model.Experience) error {

	err := p.connectionPool.Create(&experience).Error

	if err != nil {
        panic(err)
	}

	log.Printf("\ndisplay the ids of the newly inserted objects: %v", experience.ID)

	return err
}
func (p *PostgreSqlExperienceRepository) EditExperience(experience *model.Experience) error {

	err := p.connectionPool.Save(&experience).Error

	if err != nil {
        panic(err)
	}

	log.Printf("\ndisplay the ids of the edited objects: %v", experience.ID)

	return err
}
func (p *PostgreSqlExperienceRepository) DeleteExperience(id int) error {

	err := p.connectionPool.Delete(&model.Experience{}, id).Error

	if err != nil {
        panic(err)
	}

	log.Println("deleting the first result from the search filter\n"+ "The id of the deleted document:"+strconv.Itoa(id))

	return err
}