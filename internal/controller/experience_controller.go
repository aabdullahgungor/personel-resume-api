package controller

import (
	"errors"
	"net/http"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/aabdullahgungor/personal-resume-api/internal/service/interfaces"
	"github.com/gin-gonic/gin"
)

type experienceController struct {
	service interfaces.IExperienceService
}

func NewExperienceController(es interfaces.IExperienceService) *experienceController {
	return &experienceController{
		service: es,
	}
}

func (es *experienceController) GetAllExperiences(context *gin.Context) {
	experiences, err := es.service.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, experiences)
}
func (es *experienceController) GetExperienceById(context *gin.Context) {
	str_id := context.Param("id")
	experience, err := es.service.GetById(str_id)
	if err != nil {
		if errors.Is(err, service.ErrExperienceIDIsNotValid) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, service.ErrExperienceNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, experience)
}
func (es *experienceController) CreateExperience(context *gin.Context) {
	var experience model.Experience
	err := context.ShouldBindJSON(&experience)

	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = es.service.Create(&experience)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create experience: " + err.Error(),
		})
		return
	}
	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Experience has been created"})
}
func (es *experienceController) EditExperience(context *gin.Context) {

	var experience model.Experience
	err := context.ShouldBindJSON(&experience)

	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = es.service.Edit(&experience)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit experience: " + err.Error(),
		})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Experience has been edited", "experience_id": experience.ID})
}
func (es *experienceController) DeleteExperience(context *gin.Context) {
	str_id := context.Param("id")
	err := es.service.Delete(str_id)
	if err != nil {
		if errors.Is(err, service.ErrExperienceIDIsNotValid) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, service.ErrExperienceNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusAccepted, gin.H{"message": "Experience has been deleted", "experience_id": str_id})
}
