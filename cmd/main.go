package main

import (
	"github.com/aabdullahgungor/personal-resume-api/internal/database"
	"github.com/aabdullahgungor/personal-resume-api/internal/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()
	s.Run()

}