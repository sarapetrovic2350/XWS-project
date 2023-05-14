package service

import (
	reservation "common/proto/reservation-service/pb"
	user "common/proto/user-service/pb"
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"user-service/model"
	"user-service/repository"
)

type UserService struct {
	// NoSQL: injecting user repository
	UserRepo                 model.UserStore
	ReservationClientAddress string
}

func NewUserService(r model.UserStore, rca string) *UserService {
	return &UserService{
		UserRepo:                 r,
		ReservationClientAddress: rca,
	}
}

func (service *UserService) CreateUser(user *model.User) error {
	existingUser, _ := service.FindUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}
	err := service.UserRepo.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) Login(email string, password string) (string, error) {
	user, err := service.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", errors.New("incorrect password")
	}
	token, err := GenerateJWT(user.Email, user.Role, user.Id.Hex())
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

func GenerateJWT(email string, role string, id string) (string, error) {
	var sampleSecretKey = []byte("SecretYouShouldHide")
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["role"] = role
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour).Unix() * 1000
	claims["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(sampleSecretKey)
}

func (service *UserService) Delete(request *user.DeleteUserRequest) error {
	fmt.Println("In Delete User service")
	fmt.Println(request)
	fmt.Println(request.Id)
	reservationClient := repository.NewReservationClient(service.ReservationClientAddress)
	fmt.Println("reservation client created")
	getReservationsByUserIdRequest := reservation.GetUserReservationsRequest{Id: request.Id}
	reservationsResponse, err := reservationClient.GetReservationsByUserId(context.TODO(), &getReservationsByUserIdRequest)
	fmt.Println(reservationsResponse)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return service.UserRepo.Delete(request.Id)
}
func (service *UserService) Get(id primitive.ObjectID) (*model.User, error) {
	return service.UserRepo.Get(id)
}
