package router

import (
	"github.com/aabdullahgungor/personal-resume-api/internal/controller"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/gin-gonic/gin"
)

func UniversityConfigRouters(router *gin.Engine) *gin.Engine {
	
	universityRepo := repository.NewPostgreSqlUniversityRepository()
	universityService := service.NewDefaultUniversityService(universityRepo)
	universityController := controller.NewUniversityController(universityService)

	main := router.Group("api/v1")
	{
		universities := main.Group("universities")
		{
			universities.GET("/", universityController.GetAllUniversities)
			universities.GET("/:id", universityController.GetUniversityById)
			universities.POST("/", universityController.CreateUniversity)
			universities.PUT("/", universityController.EditUniversity)
			universities.DELETE("/:id", universityController.DeleteUniversity)
		}
	}
	return router
}