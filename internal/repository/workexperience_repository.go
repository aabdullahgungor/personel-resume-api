package repository

import (
	"errors"
	"log"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/database"
	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"gorm.io/gorm"
)

type IWorkExperienceRepository interface {
	GetAllWorkExperiences() ([]model.WorkExperience, error)
	GetWorkExperienceById() (model.WorkExperience, error)
	CreateWorkExperience(workExperience *model.WorkExperience) error
	EditWorkExperience(workExperience *model.WorkExperience) error
	DeleteWorkExperience(id int) error
}

var (
	ErrWorkExperienceNotFound = errors.New("FromRepository - workExperience  not found")
)

type PostgreSqlWorkExperienceRepository struct {
	connectionPool *gorm.DB
}

func NewPostgreSqlWorkExperienceRepository() *PostgreSqlWorkExperienceRepository {

	db := database.GetDatabase()

	return &PostgreSqlWorkExperienceRepository{
		connectionPool: db,
	}
}
func (p *PostgreSqlWorkExperienceRepository) GetWorkExperiences() ([]model.WorkExperience, error){

	var workExperiences []model.WorkExperience
	err := p.connectionPool.Find(&workExperiences)
	if err != nil {
        return []model.WorkExperience{}, ErrWorkExperienceNotFound
	}

	return workExperiences, nil
}
func (p *PostgreSqlWorkExperienceRepository) GetWorkExperienceById(id int) (model.WorkExperience, error) {

	var workExperience model.WorkExperience
	err := p.connectionPool.First(&workExperience, id)
	if err != nil {
		return model.WorkExperience{}, ErrWorkExperienceNotFound
	}

	return workExperience, nil
}
func (p *PostgreSqlWorkExperienceRepository) CreateWorkExperience(workExperience *model.WorkExperience) error {

	err := p.connectionPool.Create(&workExperience).Error

	if err != nil {
        panic(err)
	}

	log.Printf("\ndisplay the ids of the newly inserted objects: %v", workExperience.ID)

	return err
}
func (p *PostgreSqlWorkExperienceRepository) EditWorkExperience(workExperience *model.WorkExperience) error {

	err := p.connectionPool.Save(&workExperience).Error

	if err != nil {
        panic(err)
	}

	log.Printf("\ndisplay the ids of the edited objects: %v", workExperience.ID)

	return err
}
func (p *PostgreSqlWorkExperienceRepository) DeleteWorkExperience(id int) error {

	err := p.connectionPool.Delete(&model.WorkExperience{}, id).Error

	if err != nil {
        panic(err)
	}

	log.Println("deleting the first result from the search filter\n"+ "The id of the deleted document:"+strconv.Itoa(id))

	return err
}