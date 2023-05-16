package service

import (
	accommodation "common/proto/accommodation-service/pb"
	"context"
	"fmt"
	"reservation-service/model"
	"reservation-service/repository"
	"time"
)

type ReservationService struct {
	// NoSQL: injecting reservation repository
	ReservationRepo            model.ReservationStore
	AccommodationClientAddress string
}

func NewReservationService(r model.ReservationStore, acs string) *ReservationService {
	return &ReservationService{
		ReservationRepo:            r,
		AccommodationClientAddress: acs,
	}
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
func (service *ReservationService) GetReservationsByAccommodationId(accommodationId string) (model.Reservations, error) {
	reservations, err := service.ReservationRepo.GetAll()
	var reservationsForAccommodation model.Reservations
	for _, itr := range reservations {
		if itr.AccommodationId == accommodationId {
			reservationsForAccommodation = append(reservationsForAccommodation, itr)
		}
	}
	if err != nil {
		return nil, err
	}
	return reservationsForAccommodation, nil
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
func (service *ReservationService) GetReservationsByUserId(userId string) (model.Reservations, error) {
	fmt.Println(userId)
	fmt.Println("get active reservations by guest in reservation-service")
	reservations, err := service.ReservationRepo.GetReservationsByUserId(userId)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (service *ReservationService) GetActiveReservationsByHostId(userId string) (model.Reservations, error) {
	fmt.Println(userId)
	fmt.Println("get active reservations by host in reservation-service")
	reservations, err := service.GetAllReservations()
	if err != nil {
		return nil, err
	}

	accommodationClient := repository.NewAccommodationClient(service.AccommodationClientAddress)
	fmt.Println("accommodation client created")
	var activeReservations model.Reservations
	for _, itr := range reservations {
		getAccommodationByIdRequest := accommodation.GetAccommodationByIdRequest{Id: itr.AccommodationId}
		accommodationInReservation, _ := accommodationClient.GetAccommodationById(context.TODO(), &getAccommodationByIdRequest)
		fmt.Println(accommodationInReservation)
		if accommodationInReservation == nil {
			continue
		}
		if accommodationInReservation.Accommodation.HostID == userId {
			if itr.ReservationStatus == 1 && itr.StartDate.After(time.Now()) {
				activeReservations = append(activeReservations, itr)
			}
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
func (service *ReservationService) CancelReservation(id string) (*model.Reservation, error) {
	fmt.Println("Cancel Reservation in reservation service")
	reservation, err := service.GetById(id)
	fmt.Println(reservation)
	reservation.ReservationStatus = model.CANCELED
	err = service.ReservationRepo.Delete(id)
	if err != nil {
		return nil, err
	}
	err = service.ReservationRepo.Insert(reservation)
	if err != nil {
		return nil, err
	}
	return reservation, nil
}
