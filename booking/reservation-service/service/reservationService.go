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

func (service *ReservationService) GetActiveReservationsByGuestId(userId string) (model.Reservations, error) {
	fmt.Println(userId)
	fmt.Println("get active reservations by guest in reservation-service")
	reservations, err := service.ReservationRepo.GetReservationsByUserId(userId)
	if err != nil {
		return nil, err
	}
	var activeReservations model.Reservations
	for _, itr := range reservations {
		if itr.ReservationStatus == 1 {
			activeReservations = append(activeReservations, itr)
		}
	}
	return activeReservations, nil
}

func (service *ReservationService) GetById(id string) (*model.Reservation, error) {
	reservation, err := service.ReservationRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return reservation, nil
}

func (service *ReservationService) Delete(id string) error {
	err := service.ReservationRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
