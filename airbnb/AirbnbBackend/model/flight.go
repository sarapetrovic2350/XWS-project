package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DateTime           time.Time          `json:"dateTime" bson:"date_time"`
	Departure          string             `json:"departure" bson:"departure"`
	Arrival            string             `json:"arrival" bson:"arrival"`
	Price              float64            `json:"price" bson:"price"`
	TotalNumberOfSeats int                `json:"totalNumberOfSeats" bson:"total_number_of_seats"`
	AvailableSeats     int                `json:"availableSeats" bson:"available_seats"`
}

type Flights []*Flight

func (u *Flights) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Flight) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Flight) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
