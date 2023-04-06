package router

import (
	"github.com/aabdullahgungor/personal-resume-api/internal/controller"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/gin-gonic/gin"
)

func WorkExperienceConfigRouters(router *gin.Engine) *gin.Engine {
	
	experienceRepo := repository.NewPostgreSqlExperienceRepository()
	experienceService := service.NewDefaultExperienceService(experienceRepo)
	experienceController := controller.NewExperienceController(experienceService)

	main := router.Group("api/v1")
	{
		workExperiences := main.Group("experiences")
		{
			workExperiences.GET("/", experienceController.GetAllExperiences)
			workExperiences.GET("/:id", experienceController.GetExperienceById)
			workExperiences.POST("/", experienceController.CreateExperience)
			workExperiences.PUT("/", experienceController.EditExperience)
			workExperiences.DELETE("/:id", experienceController.DeleteExperience)
		}
	}
	return router
}