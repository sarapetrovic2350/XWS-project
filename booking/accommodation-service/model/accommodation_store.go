package model

import (
	accommodation "common/proto/accommodation-service/pb"
)

type AccommodationStore interface {
	Insert(user *Accommodation) error
	GetAll() (Accommodations, error)
	FindAccommodationByEmail(email string) (*Accommodation, error)
	GetById(id string) (*Accommodation, error)
	DeleteAll()
	SearchAccommodation(searchCriteria *accommodation.GetAccommodationsByParamsRequest) Accommodations
	AddAvailabilityForAccommodation(request *accommodation.CreateAvailabilityRequest) error
}
