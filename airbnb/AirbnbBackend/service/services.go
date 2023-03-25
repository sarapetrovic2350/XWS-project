package service

import (
	"Rest/repository"
)

type Services struct {
	UserService *UserService
}

// InitServices should be called in main.go
func InitServices(r *repository.Repositories) *Services {
	userService := NewUserService(r.UserRepo)
	return &Services{userService}
}
