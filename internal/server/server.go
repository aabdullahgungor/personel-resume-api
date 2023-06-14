package server

import (
	"log"

	"github.com/aabdullahgungor/personal-resume-api/internal/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	//gin.SetMode(gin.ReleaseMode)
	router := router.ConfigRouters(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run("localhost:" + s.port))

}
