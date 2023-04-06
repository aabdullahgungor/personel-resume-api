package router

import (
	"github.com/aabdullahgungor/personal-resume-api/internal/controller"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/gin-gonic/gin"
)

func ConfigRouters(router *gin.Engine) *gin.Engine {

	// Abilitiy 
	abilityRepo := repository.NewPostgreSqlAbilityRepository()
	abilityService := service.NewDefaultAbilityService(abilityRepo)
	abilityController := controller.NewAbilityController(abilityService)

	// Experience
	experienceRepo := repository.NewPostgreSqlExperienceRepository()
	experienceService := service.NewDefaultExperienceService(experienceRepo)
	experienceController := controller.NewExperienceController(experienceService)

	// Personal
	personalRepo := repository.NewPostgreSqlPersonalRepository()
	personalService := service.NewDefaultPersonalService(personalRepo)
	personalController := controller.NewPersonalController(personalService)

	//University
	universityRepo := repository.NewPostgreSqlUniversityRepository()
	universityService := service.NewDefaultUniversityService(universityRepo)
	universityController := controller.NewUniversityController(universityService)

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

		experiences := main.Group("experiences")
		{
			experiences.GET("/", experienceController.GetAllExperiences)
			experiences.GET("/:id", experienceController.GetExperienceById)
			experiences.POST("/", experienceController.CreateExperience)
			experiences.PUT("/", experienceController.EditExperience)
			experiences.DELETE("/:id", experienceController.DeleteExperience)
		}

		personals := main.Group("personals")
		{
			personals.GET("/", personalController.GetAllPersonals)
			personals.GET("/:id", personalController.GetPersonalById)
			personals.POST("/", personalController.CreatePersonal)
			personals.PUT("/", personalController.EditPersonal)
			personals.DELETE("/:id", personalController.DeletePersonal)
		}

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