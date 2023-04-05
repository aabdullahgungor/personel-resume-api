package router

import (
	"github.com/aabdullahgungor/personal-resume-api/internal/controller"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/gin-gonic/gin"
)

func PersonalConfigRouters(router *gin.Engine) *gin.Engine {
	
	personalRepo := repository.NewPostgreSqlPersonalRepository()
	personalService := service.NewDefaultPersonalService(personalRepo)
	personalController := controller.NewPersonalController(personalService)

	main := router.Group("api/v1")
	{
		personals := main.Group("personals")
		{
			personals.GET("/", personalController.GetAllPersonals)
			personals.GET("/:id", personalController.GetPersonalById)
			personals.POST("/", personalController.CreatePersonal)
			personals.PUT("/", personalController.EditPersonal)
			personals.DELETE("/:id", personalController.DeletePersonal)
		}
	}
	return router
}