package model

import (
	accommodation "common/proto/accommodation-service/pb"
)

type AccommodationStore interface {
	Insert(user *Accommodation) error
	GetAll() (Accommodations, error)
	GetById(id string) (*Accommodation, error)
	DeleteAll()
	Delete(id string) error
	SearchAccommodation(searchCriteria *accommodation.GetAccommodationsByParamsRequest) Accommodations
	AddAvailabilityForAccommodation(accommodation2 *Accommodation, availability *Availability) error
}
