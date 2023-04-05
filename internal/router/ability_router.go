package router

import (
	"github.com/aabdullahgungor/personal-resume-api/internal/controller"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/gin-gonic/gin"
)

func AbilityConfigRouters(router *gin.Engine) *gin.Engine {
	
	abilityRepo := repository.NewPostgreSqlAbilityRepository()
	abilityService := service.NewDefaultAbilityService(abilityRepo)
	abilityController := controller.NewAbilityController(abilityService)

	main := router.Group("api/v1")
	{
		abilities := main.Group("abilities")
		{
			abilities.GET("/", abilityController.GetAllAbilities)
			abilities.GET("/:id", abilityController.GetAbilityById)
			abilities.POST("/", abilityController.CreateAbility)
			abilities.PUT("/", abilityController.EditAbility)
			abilities.DELETE("/:id", abilityController.DeleteAbility)
		}
	}
	return router
}