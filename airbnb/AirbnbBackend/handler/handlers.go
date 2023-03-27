package handler

import (
	"Rest/service"
	"log"
)

type Handlers struct {
	UserHandler   *UserHandler
	FlightHandler *FlightHandler
}

// InitHandlers should be called in main.go
func InitHandlers(l *log.Logger, s *service.Services) *Handlers {
	userHandler := NewUserHandler(l, s.UserService)
	flightHandler := NewFlightHandler(l, s.FlightService)
	return &Handlers{userHandler, flightHandler}
}
