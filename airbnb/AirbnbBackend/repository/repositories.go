package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Repositories struct {
	UsersRepo   *UsersRepo
	PatientRepo *PatientRepo
}

// InitRepositories should be called in main.go
func InitRepositories(client *mongo.Client, logger *log.Logger) *Repositories {
	userRepo, _ := NewUsersRepo(client, logger)
	patientRepo, _ := NewPatientRepo(client, logger)
	return &Repositories{
		UsersRepo:   userRepo,
		PatientRepo: patientRepo,
	}
}
