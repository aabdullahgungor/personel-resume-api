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
	ErrUniversityNotFound = errors.New("FromRepository - university not found")
)

type PostgreSqlUniversityRepository struct {
	connectionPool *gorm.DB
}

func NewPostgreSqlUniversityRepository() *PostgreSqlUniversityRepository {

	db := database.GetDatabase()

	return &PostgreSqlUniversityRepository{
		connectionPool: db,
	}
}

func (p *PostgreSqlUniversityRepository) GetAllUniversities() ([]model.University, error) {

	var universities []model.University
	result := p.connectionPool.Find(&universities)
	if result.Error != nil {
		return []model.University{}, ErrUniversityNotFound
	}

	return universities, nil
}
func (p *PostgreSqlUniversityRepository) GetUniversityById(id int) (model.University, error) {

	var university model.University
	result := p.connectionPool.First(&university, id)
	if result.Error != nil {
		return model.University{}, ErrUniversityNotFound
	}

	return university, nil
}
func (p *PostgreSqlUniversityRepository) CreateUniversity(university *model.University) error {

	err := p.connectionPool.Create(&university).Error

	if err != nil {
		panic(err)
	}

	log.Printf("\ndisplay the ids of the newly inserted objects: %v", university.ID)

	return err
}
func (p *PostgreSqlUniversityRepository) EditUniversity(university *model.University) error {

	err := p.connectionPool.Save(&university).Error

	if err != nil {
		panic(err)
	}

	log.Printf("\ndisplay the ids of the edited objects: %v", university.ID)

	return err
}
func (p *PostgreSqlUniversityRepository) DeleteUniversity(id int) error {

	err := p.connectionPool.Delete(&model.University{}, id).Error

	if err != nil {
		panic(err)
	}

	log.Println("deleting the first result from the search filter\n" + "The id of the deleted document:" + strconv.Itoa(id))

	return err
}
