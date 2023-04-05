package server

import (
	"log"

	"github.com/aabdullahgungor/personal-resume-api/internal/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port: "8000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	abilityRouter := router.AbilityConfigRouters(s.server)
	personalRouter := router.PersonalConfigRouters(s.server)
	universityRouter := router.UniversityConfigRouters(s.server)
	workExperienceRouter := router.WorkExperienceConfigRouters(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(abilityRouter.Run("localhost:" + s.port))
	log.Fatal(personalRouter.Run("localhost:" + s.port))
	log.Fatal(universityRouter.Run("localhost:" + s.port))
	log.Fatal(workExperienceRouter.Run("localhost:" + s.port))
}