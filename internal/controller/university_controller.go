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

// GetUniversities           godoc
// @Summary		Get universities array
// @Description	Responds with the list of all universities as JSON.
// @Tags			universities
// @Produce		json
// @Success		200	{object}	model.University
// @Router			/universities [get]
func (us *universityController) GetAllUniversities(context *gin.Context) {
	universities, err := us.service.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, universities)
}

// GetUniversity          godoc
// @Summary		Get single university  by id
// @Description	Returns the university  whose id value matches the id.
// @Tags			universities
// @Produce		json
// @Param			id path	string true "search university  by id"
// @Success		200		{object}	model.University
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/universities/{id} [get]
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

// CreateUniversity           godoc
// @Summary		Add a new university
// @Description	Takes a university  JSON and store in DB. Return saved JSON.
// @Tags			universities
// @Produce		json
// @Param			university  body	model.University 	true "Ability JSON"
// @Success		200		{object}	model.University
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/universities  [post]
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

// EditUniversity           godoc
// @Summary		Edit an university
// @Description	Takes a university JSON and edit an in DB. Return saved JSON.
// @Tags			universities
// @Produce		json
// @Param			university body	model.University	true "University JSON"
// @Success		200		{object}	model.University
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/universities [put]
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

// DeleteUniversity         godoc
// @Summary		Delete an university
// @Description	Remove an university from DB by id.
// @Tags			universities
// @Produce		json
// @Param			id path	string true "delete university by id"
// @Success		200		{object}	model.University
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/universities/{id} [delete]
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
