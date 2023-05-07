package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"time"
)

type Availability struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StartDate       time.Time          `json:"startDate" bson:"startDate"`
	EndDate         time.Time          `json:"endDate" bson:"endDate"`
	AccommodationId string             `json:"accommodationId" bson:"accommodationId"`
	Price           float64            `json:"price" bson:"price"`
	PriceSelection  PriceSelection     `json:"priceSelection" bson:"priceSelection"`
}

type PriceSelection int

const (
	PER_PERSON PriceSelection = iota
	PER_ACCOMMODATION
)

type Availabilities []*Availability

func (u *Availabilities) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Availability) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Availability) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
