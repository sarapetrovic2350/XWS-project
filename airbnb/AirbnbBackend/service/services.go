package service

import (
	"Rest/repository"
)

type Services struct {
	UserService   *UserService
	FlightService *FlightService
	TicketService *TicketService
}

// InitServices should be called in main.go
func InitServices(r *repository.Repositories) *Services {
	userService := NewUserService(r.UserRepo)
	flightService := NewFlightService(r.FlightRepo)
	ticketService := NewTicketService(r.TicketRepo, r.FlightRepo)
	return &Services{userService, flightService, ticketService}
}
