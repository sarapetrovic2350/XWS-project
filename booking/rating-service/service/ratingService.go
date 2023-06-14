package service

import (
	accommodation "common/proto/accommodation-service/pb"
	reservation "common/proto/reservation-service/pb"
	"context"
	"errors"
	"fmt"
	"rating-service/model"
	"rating-service/repository"
	"time"
)

type RatingService struct {
	// NoSQL: injecting rating repository
	RatingRepo                 model.RatingStore
	ReservationClientAddress   string
	AccommodationClientAddress string
}

func NewRatingService(r model.RatingStore, rca string, aca string) *RatingService {
	return &RatingService{
		RatingRepo:                 r,
		ReservationClientAddress:   rca,
		AccommodationClientAddress: aca,
	}
}

func (service *RatingService) GetAllRatingsHost() (model.RatingsHost, error) {
	users, err := service.RatingRepo.GetAllRatingsHost()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (service *RatingService) GetRatingHostById(id string) (*model.RatingHost, error) {
	ratingHost, err := service.RatingRepo.GetRatingHostById(id)
	if err != nil {
		return nil, err
	}
	return ratingHost, nil
}

func (service *RatingService) GetRatingAccommodationById(id string) (*model.RatingAccommodation, error) {
	ratingAccommodation, err := service.RatingRepo.GetRatingAccommodationById(id)
	if err != nil {
		return nil, err
	}
	return ratingAccommodation, nil
}

func (service *RatingService) CreateRatingForHost(ratingHost *model.RatingHost) (*model.RatingHost, error) {
	areValidPastReservationsForGuest := service.CheckPastReservationsForGuest(ratingHost.HostId, ratingHost.GuestId)
	if areValidPastReservationsForGuest {
		ratingHost.Date = time.Now()
		err := service.RatingRepo.InsertRatingHost(ratingHost)
		if err != nil {
			return nil, err
		}
		return ratingHost, nil
	}
	return nil, errors.New("guest does not have a reservation in past that is not canceled")

}
func (service *RatingService) CreateRatingForAccommodation(ratingAccommodation *model.RatingAccommodation) (*model.RatingAccommodation, error) {
	isGuestStayedAtAccommodation := service.CheckIfGuestStayedAtAccommodation(ratingAccommodation.AccommodationId, ratingAccommodation.GuestId)
	if isGuestStayedAtAccommodation {
		ratingAccommodation.Date = time.Now()
		fmt.Println(ratingAccommodation)
		err := service.RatingRepo.InsertRatingAccommodation(ratingAccommodation)
		if err != nil {
			return nil, err
		}
		return ratingAccommodation, nil
	}
	return nil, errors.New("guest did not stay at accommodation he wants to rate")
}
func (service *RatingService) CheckPastReservationsForGuest(hostId string, guestId string) bool {
	reservationClient := repository.NewReservationClient(service.ReservationClientAddress)
	fmt.Println("reservation client created")
	getReservationsByGuestIdRequest := reservation.GetReservationsByUserIdRequest{Id: guestId}
	reservationsForGuest, _ := reservationClient.GetReservationsByUserId(context.TODO(), &getReservationsByGuestIdRequest)
	fmt.Println(reservationsForGuest.Reservations)
	accommodationClient := repository.NewAccommodationClient(service.AccommodationClientAddress)
	fmt.Println("accommodation client created")
	for _, itr := range reservationsForGuest.Reservations {
		getAccommodationByIdRequest := accommodation.GetAccommodationByIdRequest{Id: itr.AccommodationId}
		accommodationInReservation, _ := accommodationClient.GetAccommodationById(context.TODO(), &getAccommodationByIdRequest)
		fmt.Println(accommodationInReservation)
		if accommodationInReservation == nil {
			continue
		}
		endDate, _ := time.Parse("2006-02-01", itr.EndDate)
		if endDate.Before(time.Now()) && accommodationInReservation.Accommodation.HostID == hostId && itr.ReservationStatus == 1 {
			return true
		}
	}
	return false
}
func (service *RatingService) CheckIfGuestStayedAtAccommodation(accommodationId string, guestId string) bool {
	reservationClient := repository.NewReservationClient(service.ReservationClientAddress)
	fmt.Println("reservation client created")
	getReservationsByGuestIdRequest := reservation.GetReservationsByUserIdRequest{Id: guestId}
	reservationsForGuest, _ := reservationClient.GetReservationsByUserId(context.TODO(), &getReservationsByGuestIdRequest)
	fmt.Println(reservationsForGuest.Reservations)
	accommodationClient := repository.NewAccommodationClient(service.AccommodationClientAddress)
	fmt.Println("accommodation client created")
	for _, itr := range reservationsForGuest.Reservations {
		getAccommodationByIdRequest := accommodation.GetAccommodationByIdRequest{Id: itr.AccommodationId}
		accommodationInReservation, _ := accommodationClient.GetAccommodationById(context.TODO(), &getAccommodationByIdRequest)
		fmt.Println(accommodationInReservation)
		if accommodationInReservation == nil {
			continue
		}
		endDate, _ := time.Parse("2006-02-01", itr.EndDate)
		if endDate.Before(time.Now()) && accommodationInReservation.Accommodation.Id == accommodationId && itr.ReservationStatus == 1 {
			return true
		}
	}
	return false
}
func (service *RatingService) DeleteRatingForHost(id string) error {
	err := service.RatingRepo.DeleteRatingForHost(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *RatingService) DeleteRatingForAccommodation(id string) error {
	err := service.RatingRepo.DeleteRatingForAccommodation(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *RatingService) UpdateRatingForHost(ratingHost *model.RatingHost) (*model.RatingHost, error) {
	fmt.Println("UpdateRatingForHost service")
	oldRatingForHost, err := service.GetRatingHostById(ratingHost.Id.Hex())
	fmt.Println(ratingHost)
	err = service.RatingRepo.DeleteRatingForHost(oldRatingForHost.Id.Hex())
	if err != nil {
		return nil, err
	}
	ratingHost.Date = time.Now()
	err = service.RatingRepo.InsertRatingHost(ratingHost)
	if err != nil {
		return nil, err
	}
	return ratingHost, nil
}

func (service *RatingService) UpdateRatingForAccommodation(ratingAccommodation *model.RatingAccommodation) (*model.RatingAccommodation, error) {
	fmt.Println("UpdateRatingForHost service")
	oldRatingForHost, err := service.GetRatingAccommodationById(ratingAccommodation.Id.Hex())
	fmt.Println(ratingAccommodation)
	err = service.RatingRepo.DeleteRatingForAccommodation(oldRatingForHost.Id.Hex())
	if err != nil {
		return nil, err
	}
	ratingAccommodation.Date = time.Now()
	err = service.RatingRepo.InsertRatingAccommodation(ratingAccommodation)
	if err != nil {
		return nil, err
	}
	return ratingAccommodation, nil
}
