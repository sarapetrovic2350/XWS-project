package service

import (
	accommodation "common/proto/accommodation-service/pb"
	reservation "common/proto/reservation-service/pb"
	user "common/proto/user-service/pb"
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
	UserClientAddress          string
}

func NewRatingService(r model.RatingStore, rca string, aca string, uca string) *RatingService {
	return &RatingService{
		RatingRepo:                 r,
		ReservationClientAddress:   rca,
		AccommodationClientAddress: aca,
		UserClientAddress:          uca,
	}
}

func (service *RatingService) GetAllRatingsHost() (model.RatingsHost, error) {
	ratingsHost, err := service.RatingRepo.GetAllRatingsHost()
	if err != nil {
		return nil, err
	}
	return ratingsHost, nil
}
func (service *RatingService) GetAllRatingsAccommodation() (model.RatingsAccommodation, error) {
	ratingsAccommodation, err := service.RatingRepo.GetAllRatingsAccommodation()
	if err != nil {
		return nil, err
	}
	return ratingsAccommodation, nil
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
func (service *RatingService) GetAllRatingHostByHostId(id string) (model.RatingsHost, error) {
	fmt.Println("in GetAllRatingHostByHostId service")
	ratings, err := service.RatingRepo.GetAllRatingsHost()
	if err != nil {
		return nil, err
	}
	var ratingsForHost model.RatingsHost
	for _, itr := range ratings {
		if itr.HostId == id {
			ratingsForHost = append(ratingsForHost, itr)
		}
	}
	return ratingsForHost, nil
}
func (service *RatingService) GetAllRatingHostByGuestId(id string) (model.RatingsHost, error) {
	fmt.Println("in GetAllRatingHostByGuestId service")
	ratings, err := service.RatingRepo.GetAllRatingsHost()
	if err != nil {
		return nil, err
	}
	var ratingsByGuest model.RatingsHost
	for _, itr := range ratings {
		if itr.GuestId == id {
			ratingsByGuest = append(ratingsByGuest, itr)
		}
	}
	return ratingsByGuest, nil
}
func (service *RatingService) GetAllRatingAccommodationByGuestId(id string) (model.RatingsAccommodation, error) {
	fmt.Println("in GetAllRatingAccommodationByGuestId service")
	ratings, err := service.RatingRepo.GetAllRatingsAccommodation()
	if err != nil {
		return nil, err
	}
	var ratingsByGuest model.RatingsAccommodation
	for _, itr := range ratings {
		if itr.GuestId == id {
			ratingsByGuest = append(ratingsByGuest, itr)
		}
	}
	return ratingsByGuest, nil
}
func (service *RatingService) GetAllRatingAccommodationByHostId(hostId string) (model.RatingsAccommodation, error) {
	fmt.Println("in GetAllRatingAccommodationByHostId service")
	ratingsAccommodation, err := service.RatingRepo.GetAllRatingsAccommodation()
	if err != nil {
		return nil, err
	}
	accommodationClient := repository.NewAccommodationClient(service.AccommodationClientAddress)
	fmt.Println("accommodation client created")
	var ratingsAccommodationByHost model.RatingsAccommodation
	for _, itr := range ratingsAccommodation {
		getAccommodationByIdRequest := accommodation.GetAccommodationByIdRequest{Id: itr.AccommodationId}
		accommodationInRating, _ := accommodationClient.GetAccommodationById(context.TODO(), &getAccommodationByIdRequest)
		fmt.Println(accommodationInRating)
		if accommodationInRating == nil {
			continue
		}
		if accommodationInRating.Accommodation.HostID == hostId {
			ratingsAccommodationByHost = append(ratingsAccommodationByHost, itr)
		}
	}
	return ratingsAccommodationByHost, nil
}

func (service *RatingService) CreateRatingForHost(ratingHost *model.RatingHost) (*model.RatingHost, error) {
	fmt.Println("usao u CreateRatingForHost")
	areValidPastReservationsForGuest := service.CheckPastReservationsForGuest(ratingHost.HostId, ratingHost.GuestId)
	if areValidPastReservationsForGuest {
		ratingHost.Date = time.Now()
		err := service.RatingRepo.InsertRatingHost(ratingHost)
		if err != nil {
			return nil, err
		}

		avgRate, _ := service.GetAvgRatingForHost(ratingHost.HostId)

		userClient := repository.NewUserClient(service.UserClientAddress)
		fmt.Println("user client created")

		accommodationClient := repository.NewAccommodationClient(service.AccommodationClientAddress)
		fmt.Println("accommodation client created")

		getUserByIdRequest := user.GetUserByIdRequest{Id: ratingHost.HostId}
		getUserByIdResponse, err := userClient.GetUserById(context.TODO(), &getUserByIdRequest)
		fmt.Println(err)
		fmt.Println("Pronalazi usera")
		fmt.Println(getUserByIdResponse)

		if avgRate > 4.7 {
			fmt.Println("Usao u if da je ocena veca od 4.7 ")
			getIfHostIsSuperHostRequest := user.GetIfHostIsSuperHostRequest{Id: ratingHost.HostId}
			getIfHostIsSuperHostResponse, err := userClient.GetIfHostIsSuperHost(context.TODO(), &getIfHostIsSuperHostRequest)
			fmt.Println(err)
			fmt.Println(getIfHostIsSuperHostResponse)
			if getIfHostIsSuperHostResponse.IsSuperHost == true {
				if getUserByIdResponse.User.IsSuperHost != true {
					getUserByIdResponse.User.IsSuperHost = true
					userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: getUserByIdResponse.User})

					getAccommodationsByHostIdRequest := accommodation.GetAccommodationsByHostIdRequest{HostId: ratingHost.HostId}
					getAccommodationsByHostIdResponse, err := accommodationClient.GetAccommodationsByHostId(context.TODO(), &getAccommodationsByHostIdRequest)
					fmt.Println(err)
					for _, itr := range getAccommodationsByHostIdResponse.Accommodations {
						itr.IsSuperHost = true
						updateAccommodationRequest := accommodation.UpdateAccommodationRequest{Accommodation: itr}
						accommodationClient.UpdateAccommodation(context.TODO(), &updateAccommodationRequest)
						fmt.Println(err)
					}
				}
			} else {
				if getUserByIdResponse.User.IsSuperHost != false {
					getUserByIdResponse.User.IsSuperHost = false
					userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: getUserByIdResponse.User})
				}
			}
		} else {
			fmt.Println("Usao u if da je ocena manja od 4.7 ")
			if getUserByIdResponse.User.IsSuperHost != false {
				getUserByIdResponse.User.IsSuperHost = false
				userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{User: getUserByIdResponse.User})
			}
		}

		return ratingHost, nil
	}

	return nil, errors.New("guest does not have a reservation in past that is not canceled")

}

func (service *RatingService) GetAvgRatingForHost(id string) (float32, error) {
	ratingsForHost, err := service.GetAllRatingHostByHostId(id)
	if err != nil {
		return 0.0, err
	}
	var totalRating uint32
	for _, rating := range ratingsForHost {
		totalRating += rating.Rate
	}
	avgRating := float32(totalRating) / float32(len(ratingsForHost))
	return avgRating, nil
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
