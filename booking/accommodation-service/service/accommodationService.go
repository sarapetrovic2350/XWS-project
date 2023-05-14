package service

import (
	"accommodation-service/model"
	accommodation "common/proto/accommodation-service/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (service *AccommodationService) FindAccommodationsByEmail(email string) (*model.Accommodation, error) {
	user, err := service.AccommodationRepo.FindAccommodationByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//func (service *AccommodationService) SearchAccommodation(searchAccommodations dto.SearchDTO) model.Accommodations {
//	accommodations := service.AccommodationRepo.SearchAccommodation(searchAccommodations)
//	var retAccommodations model.Accommodations
//	var availabilities model.Availabilities
//	for _, itr := range accommodations {
//		availabilities, _ = service.AvailabilityRepo.FindAvailabilitiesByAccommodationId(itr.Id.Hex())
//		for _, availability := range availabilities {
//			if (searchAccommodations.StartDate == availability.StartDate || searchAccommodations.StartDate.After(availability.StartDate)) &&
//				(searchAccommodations.EndDate == availability.EndDate || searchAccommodations.EndDate.Before(availability.EndDate)) {
//				if itr.MinNumberOfGuests <= searchAccommodations.NumberOfGuests && itr.MaxNumberOfGuests >= searchAccommodations.NumberOfGuests {
//					retAccommodations = append(retAccommodations, itr)
//				}
//
//			}
//
//		}
//
//	}
//	if retAccommodations != nil {
//		return retAccommodations
//	}
//	return nil
//}

// dobavljanje smestaja po hostId-u, vlasnici smestaja
func (service *AccommodationService) GetAccommodationByHostId(hostId string) (model.Accommodations, error) {
	//user, _ = service.UserRepo.GetById(userId)
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

func (service *AccommodationService) AddAvailabilityForAccommodation(accommodation *model.Accommodation, availability *model.Availability) error {
	err := service.AccommodationRepo.AddAvailabilityForAccommodation(accommodation.Id, availability)
	if err != nil {
		return err
	}
	return nil
}

func (service AccommodationService) SearchAccommodation(searchAccommodations accommodation.GetAccommodationsByParamsRequest) model.Accommodations {
	accommodations := service.AccommodationRepo.SearchAccommodation(searchAccommodations)
	var retAccommodations model.Accommodations
	for _, itr := range accommodations {
		for _, availability := range itr.Availabilities {
			if (searchAccommodations.GetSearchParams().StartDate == timestamppb.New(availability.StartDate) && searchAccommodations.GetSearchParams().StartDate.AsTime().After(availability.StartDate)) &&
				(searchAccommodations.GetSearchParams().EndDate == timestamppb.New(availability.EndDate) && searchAccommodations.GetSearchParams().EndDate.AsTime().Before(availability.EndDate)) {
				if itr.MinNumberOfGuests <= int(searchAccommodations.GetSearchParams().NumberOfGuests) && itr.MaxNumberOfGuests >= int(searchAccommodations.GetSearchParams().NumberOfGuests) {
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
