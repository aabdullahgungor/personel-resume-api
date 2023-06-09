package main

import (
	"github.com/aabdullahgungor/personal-resume-api/internal/database"
	"github.com/aabdullahgungor/personal-resume-api/internal/database/migrations"
	"github.com/aabdullahgungor/personal-resume-api/internal/server"
)

func main() {
	database.StartDB()
	db := database.GetDatabase()
	migrations.RunMigrations(db)
	s := server.NewServer()
	s.Run()

}