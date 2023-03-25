package handler

import (
	"Rest/repository"
	"log"
)

type Handlers struct {
	PatientsHandler *PatientsHandler
	UsersHandler    *UsersHandler
}

// InitHandlers should be called in main.go

func InitHandlers(l *log.Logger, r *repository.Repositories) *Handlers {
	patientsHandler := NewPatientsHandler(l, r.PatientRepo)
	usersHandler := NewUsersHandler(l, r.UsersRepo)
	return &Handlers{
		PatientsHandler: patientsHandler,
		UsersHandler:    usersHandler,
	}
}
