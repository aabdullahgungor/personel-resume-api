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

// GetExperiences           godoc
// @Summary		Get experiences   array
// @Description	Responds with the list of all experiences   as JSON.
// @Tags			experiences
// @Produce		json
// @Success		200	{object}	model.Experience
// @Router			/experiences   [get]
func (es *experienceController) GetAllExperiences(context *gin.Context) {
	experiences, err := es.service.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, experiences)
}

// GetExperience           godoc
// @Summary		Get single experience by id
// @Description	Returns the experience whose id value matches the id.
// @Tags			experiences
// @Produce		json
// @Param			id path	string true "search experience by id"
// @Success		200		{object}	model.Experience
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/experiences/{id} [get]
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

// CreateExperience           godoc
// @Summary		Add a new experience
// @Description	Takes a experience  JSON and store in DB. Return saved JSON.
// @Tags			experiences
// @Produce		json
// @Param			experience  body	model.Experience  true "Experience  JSON"
// @Success		200		{object}	model.Experience
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/experiences [post]
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

// EditExperience          godoc
// @Summary		Edit an experience
// @Description	Takes a experience  JSON and edit an in DB. Return saved JSON.
// @Tags			experiences
// @Produce		json
// @Param			experience  body	model.Experience 	true "Experience  JSON"
// @Success		200		{object}	model.Experience
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/experiences  [put]
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

// DeleteExperience           godoc
// @Summary		Delete an experience
// @Description	Remove an experience  from DB by id.
// @Tags			experiences
// @Produce		json
// @Param			id path	string true "delete experience  by id"
// @Success		200		{object}	model.Experience
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/experiences/{id} [delete]
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
