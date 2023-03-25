package service

import (
	"Rest/model"
	"Rest/repository"
)

type UserService struct {
	// NoSQL: injecting user repository
	UserRepo *repository.UserRepo
}

func NewUserService(r *repository.UserRepo) *UserService {
	return &UserService{r}
}

func (service *UserService) CreateUser(user *model.User) error {
	err := service.UserRepo.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) GetAllUsers() (model.Users, error) {
	users, err := service.UserRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
