package controller

import (
	"errors"
	"net/http"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/aabdullahgungor/personal-resume-api/internal/service/interfaces"
	"github.com/gin-gonic/gin"
)

type abilityController struct {
	service interfaces.IAbilityService
}

func NewAbilityController(as interfaces.IAbilityService) *abilityController {
	return &abilityController{
		service: as,
	}
}

// GetAbilities            godoc
// @Summary		Get abilities array
// @Description	Responds with the list of all abilities as JSON.
// @Tags			abilities
// @Produce		json
// @Success		200	{object}	model.Ability
// @Router			/abilities [get]
func (as *abilityController) GetAllAbilities(context *gin.Context) {
	abilities, err := as.service.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, abilities)
}

// GetAbility          godoc
// @Summary		Get single ability by id
// @Description	Returns the ability whose id value matches the id.
// @Tags			abilities
// @Produce		json
// @Param			id path	string true "search ability by id"
// @Success		200		{object}	model.Ability
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/abilities/{id} [get]
func (as *abilityController) GetAbilityById(context *gin.Context) {
	str_id := context.Param("id")
	ability, err := as.service.GetById(str_id)
	if err != nil {
		if errors.Is(err, service.ErrAbilityIDIsNotValid) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, service.ErrAbilityNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, ability)
}

// CreateAbility           godoc
// @Summary		Add a new ability
// @Description	Takes a ability JSON and store in DB. Return saved JSON.
// @Tags			abilities
// @Produce		json
// @Param			ability body	model.Ability	true "Ability JSON"
// @Success		200		{object}	model.Ability
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/abilities [post]
func (as *abilityController) CreateAbility(context *gin.Context) {
	var ability model.Ability
	err := context.ShouldBindJSON(&ability)

	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = as.service.Create(&ability)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create ability: " + err.Error(),
		})
		return
	}
	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Ability has been created"})
}

// EditAbility           godoc
// @Summary		Edit an ability
// @Description	Takes a ability JSON and edit an in DB. Return saved JSON.
// @Tags			abilities
// @Produce		json
// @Param			ability body	model.Ability	true "Ability JSON"
// @Success		200		{object}	model.Ability
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/abilities [put]
func (as *abilityController) EditAbility(context *gin.Context) {
	var ability model.Ability
	err := context.ShouldBindJSON(&ability)

	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = as.service.Edit(&ability)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit ability: " + err.Error(),
		})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Ability has been edited", "ability_id": ability.ID})
}

// DeleteAbility          godoc
// @Summary		Delete an ability
// @Description	Remove an ability from DB by id.
// @Tags			abilities
// @Produce		json
// @Param			id path	string true "delete ability by id"
// @Success		200		{object}	model.Ability
// @Failure 	400     error message
// @Failure 	406     error message
// @Router			/abilities/{id} [delete]
func (as *abilityController) DeleteAbility(context *gin.Context) {
	str_id := context.Param("id")
	err := as.service.Delete(str_id)
	if err != nil {
		if errors.Is(err, service.ErrAbilityIDIsNotValid) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if errors.Is(err, service.ErrAbilityNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusAccepted, gin.H{"message": "Ability has been deleted", "ability_id": str_id})
}
