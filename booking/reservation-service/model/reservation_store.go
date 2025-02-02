package model

type ReservationStore interface {
	GetAll() (Reservations, error)
	GetReservationsByUserId(id string) (Reservations, error)
	DeleteAll()
	Insert(reservation *Reservation) error
	Delete(id string) error
	GetById(id string) (*Reservation, error)
	GetReservationsByAccommodationId(id string) (Reservations, error)
}
