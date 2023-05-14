package model

import "accommodation-service/dto"

type AccommodationStore interface {
	Insert(user *Accommodation) error
	GetAll() (Accommodations, error)
	FindAccommodationByEmail(email string) (*Accommodation, error)
	GetById(id string) (*Accommodation, error)
	SearchAccommodation(searchCriteria dto.SearchDTO) Accommodations
}
