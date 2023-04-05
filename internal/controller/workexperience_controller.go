package controller

import (
	"errors"
	"net/http"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/gin-gonic/gin"
) 

type workExperienceController struct {
	service service.IWorkExperienceService
}

func NewWorkExperienceController(ws service.IWorkExperienceService) *workExperienceController {
	return &workExperienceController{
		service: ws,
	}
}

func (ws *workExperienceController) GetAllWorkExperiences(context *gin.Context) {
	workExperiences, err := ws.service.GetAll()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound,gin.H{"error": "WorkExperiences cannot show: " + err.Error(), })
		return
	}
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, workExperiences)
}
func (ws *workExperienceController) GetWorkExperienceById(context *gin.Context) {
	str_id := context.Param("id")
	workExperience, err := ws.service.GetById(str_id)
	if err != nil {
	 	if errors.Is(err, service.ErrWorkExperienceIDIsNotValid) {
	 		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is not valid"+err.Error()})
	 		return
	 	} else if  errors.Is(err, service.ErrWorkExperienceNotFound) {
	 		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "WorkExperience cannot be found"+err.Error()})
	 		return
	 	}
	 	context.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	 	return
	} 
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, workExperience)
}
func (ws *workExperienceController) CreateWorkExperience(context *gin.Context) {
	var workExperience model.WorkExperience
	err := context.ShouldBindJSON(&workExperience)

	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = ws.service.Create(&workExperience)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot create workExperience: " + err.Error(),
		})
		return
	}
	context.IndentedJSON(http.StatusCreated, gin.H{"message":"WorkExperience has been created"})
}
func (ws *workExperienceController) EditWorkExperience(context *gin.Context) {
	var workExperience model.WorkExperience
	err := context.ShouldBindJSON(&workExperience)

	if err != nil {
		context.IndentedJSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = ws.service.Edit(&workExperience)

	if err != nil {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error": "cannot edit workExperience: " + err.Error(),
		})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message":"WorkExperience has been edited","workExperience_id": workExperience.ID})
}
func (ws *workExperienceController) DeleteWorkExperience(context *gin.Context) {
	str_id := context.Param("id")
	err := ws.service.Delete(str_id)
	if err != nil {
		if errors.Is(err, service.ErrWorkExperienceIDIsNotValid) {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is not valid"+err.Error()})
			return
		} else if  errors.Is(err, service.ErrWorkExperienceNotFound) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"error": "WorkExperience cannot be found"+err.Error()})
			return
		}
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	context.IndentedJSON(http.StatusAccepted, gin.H{"message":"WorkExperience has been deleted","workExperience_id": str_id})	
}