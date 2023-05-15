package service

import (
	"accommodation-service/model"
	accommodation "common/proto/accommodation-service/pb"
	"time"
)

type AccommodationService struct {
	// NoSQL: injecting AccommodationRepo
	AccommodationRepo model.AccommodationStore
	//AvailabilityRepo  *repository.AvailabilityRepo
}

func NewAccommodationService(accommodationRepository model.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		AccommodationRepo: accommodationRepository,
	}
}

func (service *AccommodationService) CreateAccommodation(accommodation *model.Accommodation) error {
	err := service.AccommodationRepo.Insert(accommodation)
	if err != nil {
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
	newAvailabilities := append(accommodation2.Availabilities, availability)
	accommodation2.Availabilities = newAvailabilities
	accommodationObjectID := (accommodation2.Id).Hex()
	err := service.AccommodationRepo.Delete(accommodationObjectID)
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
	var retAccommodations model.Accommodations
	for _, itr := range accommodations {
		for _, availability := range itr.Availabilities {
			startDate, _ := time.Parse("2006-01-02", searchAccommodations.SearchParams.StartDate)
			endDate, _ := time.Parse("2006-01-02", searchAccommodations.SearchParams.EndDate)
			if startDate == availability.StartDate && startDate.After(availability.StartDate) &&
				endDate == availability.EndDate && endDate.Before(availability.EndDate) {
				if itr.MinNumberOfGuests <= int(searchAccommodations.SearchParams.NumberOfGuests) && itr.MaxNumberOfGuests >= int(searchAccommodations.SearchParams.NumberOfGuests) {
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
