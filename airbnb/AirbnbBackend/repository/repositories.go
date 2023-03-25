package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Repositories struct {
	UserRepo    *UserRepo
	PatientRepo *PatientRepo
}

// InitRepositories should be called in main.go
func InitRepositories(client *mongo.Client, logger *log.Logger) *Repositories {
	userRepo := NewUserRepo(client, logger)
	patientRepo := NewPatientRepo(client, logger)
	return &Repositories{userRepo, patientRepo}
}
