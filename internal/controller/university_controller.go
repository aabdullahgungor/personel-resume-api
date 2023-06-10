package controller

import (
	"errors"
	"net/http"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/aabdullahgungor/personal-resume-api/internal/service/interfaces"
	"github.com/gin-gonic/gin"
)

type universityController struct {
	service interfaces.IUniversityService
}

func NewUniversityController(us interfaces.IUniversityService) *universityController {
	return &universityController{
		service: us,
	}
}

func (us *universityController) GetAllUniversities(context *gin.Context) {
	universities, err := us.service.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, universities)
}
func (us *universityController) GetUniversityById(context *gin.Context) {
	str_id := context.Param("id")
	university, err := us.service.GetById(str_id)
	if err != nil {
		if errors.Is(err, service.ErrUniversityIDIsNotValid) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, service.ErrUniversityNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, university)
}
func (us *universityController) CreateUniversity(context *gin.Context) {

	var university model.University
	err := context.ShouldBindJSON(&university)

	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = us.service.Create(&university)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create university: " + err.Error(),
		})
		return
	}
	context.IndentedJSON(http.StatusCreated, gin.H{"message": "University has been created"})
}
func (us *universityController) EditUniversity(context *gin.Context) {
	var university model.University
	err := context.ShouldBindJSON(&university)

	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = us.service.Edit(&university)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit university: " + err.Error(),
		})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "University has been edited", "university_id": university.ID})
}
func (us *universityController) DeleteUniversity(context *gin.Context) {
	str_id := context.Param("id")
	err := us.service.Delete(str_id)
	if err != nil {
		if errors.Is(err, service.ErrUniversityIDIsNotValid) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, service.ErrUniversityNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusAccepted, gin.H{"message": "University has been deleted", "university_id": str_id})
}
