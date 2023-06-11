package controller

import (
	"errors"
	"net/http"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/aabdullahgungor/personal-resume-api/internal/service/interfaces"
	"github.com/gin-gonic/gin"
)

type personalController struct {
	service interfaces.IPersonalService
}

func NewPersonalController(ps interfaces.IPersonalService) *personalController {
	return &personalController{
		service: ps,
	}
}

// GetPersonals           godoc
// @Summary		Get personals array
// @Description	Responds with the list of all personals as JSON.
// @Tags			personals
// @Produce		json
// @Success		200	{object}	model.Personal
// @Router			/personals [get]
func (ps *personalController) GetAllPersonals(context *gin.Context) {
	personals, err := ps.service.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, personals)
}

// GetPersonal          godoc
// @Summary		Get single personal by id
// @Description	Returns the personal whose id value matches the id.
// @Tags			personals
// @Produce		json
// @Param			id path	string true "search personal by id"
// @Success		200		{object}	model.Personal
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/personals/{id} [get]
func (ps *personalController) GetPersonalById(context *gin.Context) {
	str_id := context.Param("id")
	personal, err := ps.service.GetById(str_id)
	if err != nil {
		if errors.Is(err, service.ErrPersonalIDIsNotValid) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, service.ErrPersonalNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, personal)
}

// CreatePersonal           godoc
// @Summary		Add a new personal
// @Description	Takes a personal JSON and store in DB. Return saved JSON.
// @Tags			personals
// @Produce		json
// @Param			personal body	model.Personal	true "Personal JSON"
// @Success		200		{object}	model.Personal
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/personals [post]
func (ps *personalController) CreatePersonal(context *gin.Context) {
	var personal model.Personal

	err := context.ShouldBindJSON(&personal)
	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		context.Abort()
		return
	}
	if err := personal.HashPassword(personal.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	err = ps.service.Create(&personal)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create personal: " + err.Error(),
		})
		return
	}
	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Personal has been created"})
}

// EditPersonal           godoc
// @Summary		Edit an personal
// @Description	Takes a personal JSON and edit an in DB. Return saved JSON.
// @Tags			personals
// @Produce		json
// @Param			personal body	model.Personal	true "Personal JSON"
// @Success		200		{object}	model.Personal
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/personals [put]
func (ps *personalController) EditPersonal(context *gin.Context) {
	var personal model.Personal
	err := context.ShouldBindJSON(&personal)

	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = ps.service.Edit(&personal)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit personal: " + err.Error(),
		})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Personal has been edited", "personal_id": personal.ID})
}

// DeletePersonal         godoc
// @Summary		Delete an personal
// @Description	Remove an personal from DB by id.
// @Tags			personals
// @Produce		json
// @Param			id path	string true "delete personal by id"
// @Success		200		{object}	model.Personal
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/personals/{id} [delete]
func (ps *personalController) DeletePersonal(context *gin.Context) {
	str_id := context.Param("id")
	err := ps.service.Delete(str_id)
	if err != nil {
		if errors.Is(err, service.ErrPersonalIDIsNotValid) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, service.ErrPersonalNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusAccepted, gin.H{"message": "Personal has been deleted", "personal_id": str_id})
}
