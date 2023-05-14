package model

type AvailabilityStore interface {
	Insert(availability *Availability) error
	GetAll() (Availabilities, error)
	FindAvailabilitiesByAccommodationId(id string) (Availabilities, error)
	GetById(id string) (*Availability, error)
}
