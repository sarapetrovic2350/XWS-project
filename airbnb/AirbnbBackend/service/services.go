package service

import (
	"Rest/repository"
)

type Services struct {
	UserService   *UserService
	FlightService *FlightService
}

// InitServices should be called in main.go
func InitServices(r *repository.Repositories) *Services {
	userService := NewUserService(r.UserRepo)
	flightService := NewFlightService(r.FlightRepo)
	return &Services{userService, flightService}
}
