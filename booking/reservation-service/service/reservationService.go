package service

import (
	"fmt"
	"reservation-service/model"
)

type ReservationService struct {
	// NoSQL: injecting reservation repository
	ReservationRepo model.ReservationStore
}

func NewReservationService(r model.ReservationStore) *ReservationService {
	return &ReservationService{r}
}

func (service *ReservationService) CreateReservation(reservation *model.Reservation) error {
	reservation.ReservationStatus = model.PENDING
	err := service.ReservationRepo.Insert(reservation)
	if err != nil {
		return err
	}
	return nil
}

func (service *ReservationService) GetAllReservations() (model.Reservations, error) {
	reservations, err := service.ReservationRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (service *ReservationService) GetReservationsByUserId(userId string) (model.Reservations, error) {
	fmt.Println(userId)
	fmt.Println("get resrvations by user id reservation-service")
	reservations, err := service.ReservationRepo.GetReservationsByUserId(userId)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (service *ReservationService) Delete(id string) error {
	err := service.ReservationRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
