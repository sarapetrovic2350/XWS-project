package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NumberOfGuests  int                `json:"NumberOfGuests" bson:"number_of_guests"`
	StartDate       time.Time          `json:"startDate" bson:"startDate"`
	EndDate         time.Time          `json:"endDate" bson:"endDate"`
	UserID          string             `json:"userID" bson:"userID;"`
	AccommodationID string             `json:"accommodationID" bson:"accommodationID;"`
}

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
