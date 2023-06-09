package repository

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/aabdullahgungor/personal-resume-api/internal/database"
	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"gorm.io/gorm"
)

var (
	ErrPersonalNotFound = errors.New("FromRepository - personal not found")
)

type PostgreSqlPersonalRepository struct {
	connectionPool *gorm.DB
}

func NewPostgreSqlPersonalRepository() *PostgreSqlPersonalRepository {

	db := database.GetDatabase()

	return &PostgreSqlPersonalRepository{
		connectionPool: db,
	}
}

func (p *PostgreSqlPersonalRepository) GetAllPersonals() ([]model.Personal, error) {

	var personals []model.Personal
	result := p.connectionPool.Preload("Abilities").Preload("Universities").Find(&personals)
	if result.Error != nil {
		return []model.Personal{}, ErrPersonalNotFound
	}

	return personals, nil
}
func (p *PostgreSqlPersonalRepository) GetPersonalById(id int) (model.Personal, error) {

	var personal model.Personal
	result := p.connectionPool.Preload("Abilities").Preload("Universities").First(&personal, id)
	if result.Error != nil {
		return model.Personal{}, ErrPersonalNotFound
	}

	return personal, nil
}
func (p *PostgreSqlPersonalRepository) GetPersonalByEmail(email string) (model.Personal, error) {

	var personal model.Personal
	result := p.connectionPool.Preload("Abilities").Preload("Universities").First(&personal, email)
	if result.Error != nil {
		fmt.Println("Repository - Personal Not Found")
		return model.Personal{}, ErrPersonalNotFound
	}

	return personal, nil
}
func (p *PostgreSqlPersonalRepository) CreatePersonal(personal *model.Personal) error {

	err := p.connectionPool.Create(&personal).Error

	if err != nil {
		panic(err)
	}

	log.Printf("\ndisplay the ids of the newly inserted objects: %v", personal.ID)

	return err
}
func (p *PostgreSqlPersonalRepository) EditPersonal(personal *model.Personal) error {

	err := p.connectionPool.Save(&personal).Error

	if err != nil {
		panic(err)
	}

	log.Printf("\ndisplay the ids of the edited objects: %v", personal.ID)

	return err
}
func (p *PostgreSqlPersonalRepository) DeletePersonal(id int) error {

	err := p.connectionPool.Delete(&model.Personal{}, id).Error

	if err != nil {
		panic(err)
	}

	log.Println("deleting the first result from the search filter\n" + "The id of the deleted document:" + strconv.Itoa(id))

	return err
}
