package service

import (
	"Rest/dto"
	"Rest/model"
	"Rest/repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
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

func (service *UserService) Login(dto *dto.Login) (string, error) {
	user, err := service.FindUserByEmail(dto.Email)
	if err != nil {
		return "", err
	}
	if user.Password != dto.Password {
		return "", errors.New("incorrect password")
	}
	token, err := GenerateJWT(user.Email, user.Role)
	if err != nil {
		return "", err
	}
	return token, err
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

func GenerateJWT(email string, role string) (string, error) {
	var sampleSecretKey = []byte("SecretYouShouldHide")
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour).Unix() * 1000
	claims["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(sampleSecretKey)
}

func (service *UserService) DeleteGuestUser(id string) error {
	err := service.UserRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
