package service

import (
	accommodation "common/proto/accommodation-service/pb"
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
	UserRepo                   model.UserStore
	ReservationClientAddress   string
	AccommodationClientAddress string
	RatingClientAddress        string
	orchestrator               *DeleteUserOrchestrator
}

func NewUserService(r model.UserStore, rca string, aca string, rtca string, orchestrator *DeleteUserOrchestrator) *UserService {
	return &UserService{
		UserRepo:                   r,
		ReservationClientAddress:   rca,
		AccommodationClientAddress: aca,
		RatingClientAddress:        rtca,
		orchestrator:               orchestrator,
	}
}

func (service *UserService) CreateUser(user *model.User) error {
	existingUser, _ := service.FindUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}
	user.IsSuperHost = false
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
func (service *UserService) GetAllHostUsers() (model.Users, error) {
	users, err := service.UserRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var hostUsers model.Users
	for _, itr := range users {
		if itr.Role == "HOST" {
			hostUsers = append(hostUsers, itr)
		}
	}
	return hostUsers, nil
}

func (service *UserService) FindUserByEmail(email string) (*model.User, error) {
	user, err := service.UserRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (service *UserService) FindUserById(id string) (*model.User, error) {
	user, err := service.UserRepo.FindUserById(id)
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
func (service *UserService) DeleteUser(id primitive.ObjectID) error {
	fmt.Println("In DeleteUser User service")
	fmt.Println(id)
	user, err := service.Get(id)
	if err != nil {
		return err
	}
	err = service.orchestrator.Start(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (service *UserService) DeleteGuestUser(request *user.DeleteUserRequest) error {
	fmt.Println("In Delete Guest User service")
	fmt.Println(request)
	fmt.Println(request.Id)
	reservations, err := service.GetActiveReservationsForGuestUser(request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(reservations.Reservations)
	if reservations.Reservations != nil {
		return errors.New("guest user has active reservations")
	}
	return service.UserRepo.Delete(request.Id)
}

func (service *UserService) DeleteHostUser(request *user.DeleteUserRequest) error {
	fmt.Println("In Delete Host User service")
	fmt.Println(request)
	fmt.Println(request.Id)
	reservations, err := service.GetActiveReservationsForHostUser(request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(reservations.Reservations)
	if reservations.Reservations != nil {
		return errors.New("host user has active reservations")
	}
	err = service.UserRepo.Delete(request.Id)
	if err != nil {
		return errors.New("error while deleting host user")
	}
	// delete all accommodations created by host
	accommodationClient := repository.NewAccommodationClient(service.AccommodationClientAddress)
	fmt.Println("accommodation client created")
	deleteAccommodationsByHostRequest := accommodation.DeleteAccommodationsByHostIdRequest{Id: request.Id}
	_, err = accommodationClient.DeleteAccommodationsByHostId(context.TODO(), &deleteAccommodationsByHostRequest)
	if err != nil {
		return errors.New("error while deleting accommodations created by host")
	}
	return nil
}

func (service *UserService) GetActiveReservationsForGuestUser(request *user.DeleteUserRequest) (*reservation.GetActiveReservationsResponse, error) {
	fmt.Println("In GetActiveReservationsForGuestUser User service")
	fmt.Println(request)
	fmt.Println(request.Id)
	reservationClient := repository.NewReservationClient(service.ReservationClientAddress)
	fmt.Println("reservation client created")
	getReservationsByUserIdRequest := reservation.GetActiveReservationsRequest{Id: request.Id}
	reservationsResponse, err := reservationClient.GetActiveReservationsByGuestId(context.TODO(), &getReservationsByUserIdRequest)
	return reservationsResponse, err
}

func (service *UserService) GetActiveReservationsForHostUser(request *user.DeleteUserRequest) (*reservation.GetActiveReservationsResponse, error) {
	fmt.Println("In GetActiveReservationsForHostUser User service")
	fmt.Println(request)
	fmt.Println(request.Id)
	reservationClient := repository.NewReservationClient(service.ReservationClientAddress)
	fmt.Println("reservation client created")
	getReservationsByUserIdRequest := reservation.GetActiveReservationsRequest{Id: request.Id}
	reservationsResponse, err := reservationClient.GetActiveReservationsByHostId(context.TODO(), &getReservationsByUserIdRequest)
	return reservationsResponse, err
}

func (service *UserService) GetIfHostIsSuperHost(id string) (bool, error) {
	fmt.Println("In GetActiveReservationsForHostUser User service")
	fmt.Println(id)
	fmt.Println(id)
	reservationClient := repository.NewReservationClient(service.ReservationClientAddress)
	fmt.Println("reservation client created")

	// proveramo da li je imao 5 ili vise rezervacija
	getNumberOfPastReservationsByHostRequest := reservation.GetNumberOfPastReservationsByHostRequest{Id: id}
	numberOfPastReservationsByHostResponse, err := reservationClient.GetNumberOfPastReservationsByHostId(context.TODO(), &getNumberOfPastReservationsByHostRequest)

	// proveravamo da li je trajanje rezervacije trjalo 50 ili duze dana
	// treba 50 danaaaaaa
	getDurationOfPastReservationsByHostIdRequest := reservation.GetDurationOfPastReservationsByHostIdRequest{Id: id}
	durationOfPastReservationsByHostResponse, err := reservationClient.GetDurationOfPastReservationsByHostId(context.TODO(), &getDurationOfPastReservationsByHostIdRequest)

	//provera da je canceled manje od 5%
	getAcceptanceRateByHostIdRequest := reservation.GetAcceptanceRateByHostIdRequest{Id: id}
	getAcceptanceRateByHostIdResponse, err := reservationClient.GetAcceptanceRateByHostId(context.TODO(), &getAcceptanceRateByHostIdRequest)

	if numberOfPastReservationsByHostResponse.NumReservations < 5 {
		return false, err
	} else if durationOfPastReservationsByHostResponse.NumDays < 50 {
		return false, err
	} else if getAcceptanceRateByHostIdResponse.Percentage > 5 {
		return false, err
	}

	return true, err
}

func (service *UserService) Get(id primitive.ObjectID) (*model.User, error) {
	return service.UserRepo.Get(id)
}

func (service *UserService) Update(user *model.User) error {
	checkUser, err := service.UserRepo.FindUserByEmail(user.Email)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return err
	}
	fmt.Print("Get user by email: ")
	fmt.Println(checkUser)
	stringObjectID := (user.Id).Hex()
	fmt.Print(stringObjectID)
	err = service.UserRepo.Delete(stringObjectID)
	if err != nil {
		return err
	}
	err = service.UserRepo.Insert(user)
	if err != nil {
		return err
	}
	return nil
}
func (service *UserService) Delete(id string) error {
	err := service.UserRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
