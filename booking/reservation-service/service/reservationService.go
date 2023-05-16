package service

import (
	accommodation "common/proto/accommodation-service/pb"
	"context"
	"errors"
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
	reservations, err := service.ReservationRepo.GetAll()
	for _, itr := range reservations {
		if itr.AccommodationId == reservation.AccommodationId {
			//reservationsByAccomodation = append(reservationsByAccomodation, itr)
			if itr.ReservationStatus != 1 {
				if (reservation.StartDate == itr.StartDate || reservation.StartDate.After(itr.StartDate) && reservation.StartDate.Before(itr.EndDate)) ||
					(reservation.EndDate == itr.EndDate || reservation.EndDate.Before(itr.EndDate) && reservation.EndDate.After(itr.StartDate)) {
					//if itr.ReservationStatus == 1 {
					//	return errors.New("Accommodation already has reservation!")
					//}
					//poruka da ne moze da napravi
					return errors.New("Accommodation already has reservation!")
				}
				if (reservation.StartDate == itr.StartDate || reservation.StartDate.Before(itr.StartDate)) &&
					(reservation.EndDate == itr.EndDate || reservation.EndDate.After(itr.EndDate)) {
					//if itr.ReservationStatus == 1 {
					//	return errors.New("Accommodation already has reservation!")
					//}
					return errors.New("Accommodation already has reservation!")
				}
			}
		}
	}

	err = service.ReservationRepo.Insert(reservation)
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
