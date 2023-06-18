package service

import (
	"accommodation-service/model"
	"accommodation-service/repository"
	accommodation "common/proto/accommodation-service/pb"
	reservation "common/proto/reservation-service/pb"
	user "common/proto/user-service/pb"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

type AccommodationService struct {
	// NoSQL: injecting AccommodationRepo
	AccommodationRepo model.AccommodationStore
	//AvailabilityRepo  *repository.AvailabilityRepo
	ReservationClientAddress string
	UserClientAddress        string
}

func NewAccommodationService(accommodationRepository model.AccommodationStore, rca string, uca string) *AccommodationService {
	return &AccommodationService{
		AccommodationRepo:        accommodationRepository,
		ReservationClientAddress: rca,
		UserClientAddress:        uca,
	}
}

func (service *AccommodationService) CreateAccommodation(accommodation *model.Accommodation) error {
	//accommodation.IsSuperHost = false

	userClient := repository.NewUserClient(service.UserClientAddress)
	fmt.Println("user client created")

	getUserByIdRequest := user.GetUserByIdRequest{Id: accommodation.HostID}
	hostUser, err := userClient.GetUserById(context.TODO(), &getUserByIdRequest)
	fmt.Println(err)
	fmt.Println("Pronalazi usera")
	fmt.Println(hostUser)

	if hostUser.User.IsSuperHost {
		accommodation.IsSuperHost = true
	} else {
		accommodation.IsSuperHost = false
	}

	err1 := service.AccommodationRepo.Insert(accommodation)
	if err1 != nil {
		return err
	}
	return nil
}

func (service *AccommodationService) GetAllAccommodations() (model.Accommodations, error) {
	accommodations, err := service.AccommodationRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return accommodations, nil
}

// dobavljanje smestaja po hostId-u, vlasnici smestaja
func (service *AccommodationService) GetAccommodationsByHostId(hostId string) (model.Accommodations, error) {
	accommodations, err := service.AccommodationRepo.GetAll()
	var retAccommodations model.Accommodations
	for _, itr := range accommodations {
		if itr.HostID == hostId {
			retAccommodations = append(retAccommodations, itr)
		}
	}

	if err != nil {
		return nil, err
	}
	return retAccommodations, nil
}

func (service *AccommodationService) GetById(id string) (*model.Accommodation, error) {
	flight, err := service.AccommodationRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return flight, nil
}

func (service *AccommodationService) AddAvailabilityForAccommodation(accommodation2 *model.Accommodation, availability *model.Availability) error {
	fmt.Println("In AddAvailabilityForAccommodation accommodation service")
	reservationClient := repository.NewReservationClient(service.ReservationClientAddress)
	fmt.Println("reservation client created")
	idString := accommodation2.Id.Hex()
	getReservationsByAccommodationRequest := reservation.GetReservationsByAccommodationRequest{Id: idString}
	reservationsForAccommodation, err := reservationClient.GetReservationsByAccommodationId(context.TODO(), &getReservationsByAccommodationRequest)

	for _, itr := range reservationsForAccommodation.Reservations {
		if itr.ReservationStatus == 1 {
			startDate, _ := time.Parse("2006-02-01", itr.StartDate)
			endDate, _ := time.Parse("2006-02-01", itr.EndDate)
			if (availability.StartDate == startDate || startDate.Before(availability.StartDate)) &&
				(availability.EndDate == endDate || endDate.After(availability.EndDate)) {
				return errors.New("reservation exists for given period of time")
			}
			if (availability.StartDate == startDate || startDate.After(availability.StartDate) && startDate.Before(availability.EndDate)) &&
				(availability.EndDate == endDate || endDate.After(availability.EndDate)) {
				return errors.New("reservation exists for given period of time")
			}
			if (availability.StartDate == startDate || startDate.After(availability.StartDate)) &&
				(availability.EndDate == endDate || endDate.Before(availability.EndDate)) {
				return errors.New("reservation exists for given period of time")
			}
			if (availability.StartDate == startDate || startDate.Before(availability.StartDate) && endDate.After(availability.StartDate)) &&
				(availability.EndDate == endDate || endDate.Before(availability.EndDate)) {
				return errors.New("reservation exists for given period of time")
			}
			if availability.StartDate == endDate || availability.EndDate == startDate {
				return errors.New("reservation exists for given period of time")
			}
		}
	}
	newAvailabilities := append(accommodation2.Availabilities, availability)
	accommodation2.Availabilities = newAvailabilities
	accommodationObjectID := (accommodation2.Id).Hex()
	err = service.AccommodationRepo.Delete(accommodationObjectID)
	if err != nil {
		return err
	}
	err = service.AccommodationRepo.Insert(accommodation2)
	if err != nil {
		return err
	}
	return nil
}

func (service AccommodationService) SearchAccommodation(searchAccommodations *accommodation.GetAccommodationsByParamsRequest) model.Accommodations {
	accommodations := service.AccommodationRepo.SearchAccommodation(searchAccommodations)
	fmt.Println(accommodations)
	for _, itr := range accommodations {
		fmt.Println(itr.Name)
	}
	var retAccommodations model.Accommodations
	for _, itr := range accommodations {
		fmt.Println(itr.Availabilities)
		for _, availability := range itr.Availabilities {
			fmt.Println(searchAccommodations.SearchParams.StartDate)
			fmt.Println(searchAccommodations.SearchParams.EndDate)
			startDate1 := strings.Split(searchAccommodations.SearchParams.StartDate, "T")
			endDate1 := strings.Split(searchAccommodations.SearchParams.EndDate, "T")
			startDate, _ := time.Parse("2006-01-02", startDate1[0])
			endDate, _ := time.Parse("2006-01-02", endDate1[0])
			if (startDate == availability.StartDate || startDate.After(availability.StartDate)) &&
				(endDate == availability.EndDate || endDate.Before(availability.EndDate)) {
				itr.Availabilities = itr.Availabilities[:0]
				if itr.MinNumberOfGuests <= int(searchAccommodations.SearchParams.NumberOfGuests) && itr.MaxNumberOfGuests >= int(searchAccommodations.SearchParams.NumberOfGuests) {
					itr.Availabilities = append(itr.Availabilities, availability)
					retAccommodations = append(retAccommodations, itr)
				}

			}

		}

	}
	if retAccommodations != nil {
		return retAccommodations
	}
	return nil
}
func (service *AccommodationService) Delete(id string) error {
	err := service.AccommodationRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *AccommodationService) Update(acc *model.Accommodation) error {
	stringObjectID1 := (acc.Id).Hex()
	checkAcc, err := service.AccommodationRepo.GetById(stringObjectID1)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return err
	}
	fmt.Print("Get user by email: ")
	fmt.Println(checkAcc)
	stringObjectID := (acc.Id).Hex()
	fmt.Print(stringObjectID)
	err = service.AccommodationRepo.Delete(stringObjectID)
	if err != nil {
		return err
	}
	err = service.AccommodationRepo.Insert(acc)
	if err != nil {
		return err
	}
	return nil
}
