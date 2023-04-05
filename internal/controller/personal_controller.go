package controller

import (
	"errors"
	"net/http"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/gin-gonic/gin"
)

type personalController struct {
	service service.IPersonalService
}

func NewPersonalController(ps service.IPersonalService) *personalController {
	return &personalController{
		service: ps,
	}
}

func (ps *personalController) GetAllPersonals(context *gin.Context) {
	personals, err := ps.service.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound,gin.H{"error": "Personals cannot show: " + err.Error(), })
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, personals)
}
func (ps *personalController) GetPersonalById(context *gin.Context) {
	str_id := context.Param("id")
	personal, err := ps.service.GetById(str_id)
	if err != nil {
	 	if errors.Is(err, service.ErrPersonalIDIsNotValid) {
	 		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is not valid"+err.Error()})
	 		return
	 	} else if  errors.Is(err, service.ErrPersonalNotFound) {
	 		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Personal cannot be found"+err.Error()})
	 		return
	 	}
	 	context.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	 	return
	} 
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, personal)
}
func (ps *personalController) CreatePersonal(context *gin.Context) {
	var personal model.Personal
	err := context.ShouldBindJSON(&personal)

	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = ps.service.Create(&personal)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create personal: " + err.Error(),
		})
		return
	}
	context.IndentedJSON(http.StatusCreated, gin.H{"message":"Personal has been created"})
}
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

	context.IndentedJSON(http.StatusCreated, gin.H{"message":"Personal has been edited","personal_id": personal.ID})
}
func (ps *personalController) DeletePersonal(context *gin.Context) {
	str_id := context.Param("id")
	err := ps.service.Delete(str_id)
	if err != nil {
		if errors.Is(err, service.ErrPersonalIDIsNotValid) {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is not valid"+err.Error()})
			return
		} else if  errors.Is(err, service.ErrPersonalNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Personal cannot be found"+err.Error()})
			return
		}
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	context.IndentedJSON(http.StatusAccepted, gin.H{"message":"Personal has been deleted","personal_id": str_id})	
}