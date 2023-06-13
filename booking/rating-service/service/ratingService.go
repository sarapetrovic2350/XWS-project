package service

import (
	"rating-service/model"
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
