package service

import (
	"Rest/dto"
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
	user.Role = model.GuestUserRole
	err := service.UserRepo.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) Login(dto *dto.Login) (*model.User, error) {
	user, err := service.FindUserByEmail(dto.Email)
	if user != nil && user.Password == dto.Password {
		return user, nil
	}
	return nil, err
}

func (service *UserService) GetAllUsers() (model.Users, error) {
	users, err := service.UserRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *UserService) FindUserByEmail(email string) (*model.User, error) {
	user, err := service.UserRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
