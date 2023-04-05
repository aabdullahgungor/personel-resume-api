package router

import (
	"github.com/aabdullahgungor/personal-resume-api/internal/controller"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/gin-gonic/gin"
)

func WorkExperienceConfigRouters(router *gin.Engine) *gin.Engine {
	
	workExperienceRepo := repository.NewPostgreSqlWorkExperienceRepository()
	workExperienceService := service.NewDefaultWorkExperienceService(workExperienceRepo)
	workExperienceController := controller.NewWorkExperienceController(workExperienceService)

	main := router.Group("api/v1")
	{
		workExperiences := main.Group("workExperiences")
		{
			workExperiences.GET("/", workExperienceController.GetAllWorkExperiences)
			workExperiences.GET("/:id", workExperienceController.GetWorkExperienceById)
			workExperiences.POST("/", workExperienceController.CreateWorkExperience)
			workExperiences.PUT("/", workExperienceController.EditWorkExperience)
			workExperiences.DELETE("/:id", workExperienceController.DeleteWorkExperience)
		}
	}
	return router
}