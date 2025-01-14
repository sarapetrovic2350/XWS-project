package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	UserRepo    *UserRepo
	PatientRepo *PatientRepo
	FlightRepo  *FlightRepo
	TicketRepo  *TicketRepo
}

// InitRepositories should be called in main.go
func InitRepositories(client *mongo.Client, logger *log.Logger) *Repositories {
	userRepo := NewUserRepo(client, logger)
	patientRepo := NewPatientRepo(client, logger)
	flightRepo := NewFlightRepo(client, logger)
	ticketRepo := NewTicketRepo(client, logger)
	return &Repositories{userRepo, patientRepo, flightRepo, ticketRepo}
}
