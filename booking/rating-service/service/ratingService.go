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

func (service *RatingService) CreateRatingForHost(ratingHost *model.RatingHost) error {
	areValidPastReservationsForGuest := service.CheckPastReservationsForGuest(ratingHost.HostId, ratingHost.GuestId)
	if areValidPastReservationsForGuest {
		err := service.RatingRepo.InsertRatingHost(ratingHost)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("guest does not have a reservation in past that is not canceled")

}
func (service *RatingService) CreateRatingForAccommodation(ratingAccommodation *model.RatingAccommodation) error {
	fmt.Println("CreateRatingForAccommodation service")
	fmt.Println(ratingAccommodation)
	err := service.RatingRepo.InsertRatingAccommodation(ratingAccommodation)
	if err != nil {
		return err
	}
	return nil
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
func (service *RatingService) DeleteRatingForHost(id string) error {
	err := service.RatingRepo.DeleteRatingForHost(id)
	if err != nil {
		return err
	}
	return nil
}
func (service *RatingService) UpdateRatingForHost(ratingHost *model.RatingHost) error {
	fmt.Println("UpdateRatingForHost service")
	oldRatingForHost, err := service.GetRatingHostById(ratingHost.Id.Hex())
	fmt.Println(ratingHost)
	err = service.RatingRepo.DeleteRatingForHost(oldRatingForHost.Id.Hex())
	if err != nil {
		return err
	}
	err = service.RatingRepo.InsertRatingHost(ratingHost)
	if err != nil {
		return err
	}
	return nil
}
