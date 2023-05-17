package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NumberOfGuests    int                `json:"NumberOfGuests" bson:"number_of_guests"`
	StartDate         time.Time          `json:"startDate" bson:"startDate"`
	EndDate           time.Time          `json:"endDate" bson:"endDate"`
	UserId            string             `json:"userId" bson:"user_id"`
	AccommodationId   string             `json:"accommodationId" bson:"accommodation_id"`
	ReservationStatus ReservationStatus  `json:"reservationStatus" bson:"reservation_status"`
}
type ReservationStatus int

const (
	PENDING  ReservationStatus = 0
	ACCEPTED ReservationStatus = 1
	CANCELED ReservationStatus = 2
	REJECTED ReservationStatus = 3
)

type Reservations []*Reservation

func (u *Reservations) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Reservation) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Reservation) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
