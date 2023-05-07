package service

import (
	"Rest/model"
	"Rest/repository"
)

type AccommodationService struct {
	// NoSQL: injecting user repository
	AccommodationRepo *repository.AccommodationRepo
}

func NewAccommodationService(r *repository.AccommodationRepo) *AccommodationService {
	return &AccommodationService{r}
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
