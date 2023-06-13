package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type RatingHost struct {
	Id      primitive.ObjectID `bson:"_id" json:"id"`
	HostId  string             `json:"hostId" bson:"host_id"`
	GuestId string             `json:"guestId" bson:"guest_id"`
	Rate    uint32             `json:"rate" bson:"rate"`
}
type RatingAccommodation struct {
	Id              primitive.ObjectID `bson:"_id" json:"id"`
	AccommodationId string             `json:"accommodationId" bson:"accommodation_id"`
	GuestId         string             `json:"guestId" bson:"guest_id"`
	Rate            uint32             `json:"rate" bson:"rate"`
}

type RatingsHost []*RatingHost

type RatingsAccommodation []*RatingAccommodation

func (rh *RatingsHost) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(rh)
}

func (rh *RatingHost) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(rh)
}

func (rh *RatingHost) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(rh)
}

func (ra *RatingsAccommodation) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ra)
}

func (ra *RatingAccommodation) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ra)
}

func (ra *RatingAccommodation) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(ra)
}
