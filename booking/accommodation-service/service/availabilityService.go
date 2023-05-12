package service

import (
	"github.com/sarapetrovic2350/XWS-project/booking/accommodation-service/model"
	"github.com/sarapetrovic2350/XWS-project/booking/accommodation-service/repository"
)

type AvailabilityService struct {
	// NoSQL: injecting user repository
	AvailabilityRepo *repository.AvailabilityRepo
}

func NewAvailabilityService(r *repository.AvailabilityRepo) *AvailabilityService {
	return &AvailabilityService{r}
}

func (service *AvailabilityService) CreateAvailability(availability *model.Availability) error {
	err := service.AvailabilityRepo.Insert(availability)
	if err != nil {
		return err
	}
	return nil
}

func (service *AvailabilityService) GetAllAvailabilities() (model.Availabilities, error) {
	availabilities, err := service.AvailabilityRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return availabilities, nil
}

func (service *AvailabilityService) FindAvailabilitiesByAccommodationId(id string) ([]*model.Availability, error) {
	availabilities, err := service.AvailabilityRepo.FindAvailabilitiesByAccommodationId(id)
	if err != nil {
		return nil, err
	}
	return availabilities, nil
}
