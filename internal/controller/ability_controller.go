package controller

import (
	"errors"
	"net/http"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/gin-gonic/gin"
)

type abilityController struct {
	service service.IAbilityService
}

func NewAbilityController(as service.IAbilityService) *abilityController {
	return &abilityController{
		service: as,
	}
}

func (as *abilityController) GetAllAbilities(context *gin.Context) {
	abilities, err := as.service.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound,gin.H{"error": "Abilities cannot show: " + err.Error(), })
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, abilities)
}
func (as *abilityController) GetAbilityById(context *gin.Context) {
	str_id := context.Param("id")
	ability, err := as.service.GetById(str_id)
	if err != nil {
	 	if errors.Is(err, service.ErrAbilityIDIsNotValid) {
	 		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is not valid"+err.Error()})
	 		return
	 	} else if  errors.Is(err, service.ErrAbilityNotFound) {
	 		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Ability cannot be found"+err.Error()})
	 		return
	 	}
	 	context.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	 	return
	} 
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, ability)
}
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
	context.IndentedJSON(http.StatusCreated, gin.H{"message":"Ability has been created"})
}
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

	context.IndentedJSON(http.StatusCreated, gin.H{"message":"Ability has been edited","ability_id": ability.ID})
}
func (as *abilityController) DeleteAbility(context *gin.Context) {
	str_id := context.Param("id")
	err := as.service.Delete(str_id)
	if err != nil {
		if errors.Is(err, service.ErrAbilityIDIsNotValid) {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is not valid"+err.Error()})
			return
		} else if  errors.Is(err, service.ErrAbilityNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Ability cannot be found"+err.Error()})
			return
		}
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	context.IndentedJSON(http.StatusAccepted, gin.H{"message":"Ability has been deleted","ability_id": str_id})	
}