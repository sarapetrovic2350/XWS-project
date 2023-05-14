package model

type ReservationStore interface {
	GetAll() (Reservations, error)
	GetReservationsByUserId(id string) (Reservations, error)
	DeleteAll()
	Insert(reservation *Reservation) error
	Delete(id string) error
}
