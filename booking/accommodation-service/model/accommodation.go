package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Accommodation struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Availabilities    []*Availability    `json:"availabilities" bson:"availabilities"`
	Name              string             `json:"name" bson:"name"`
	MinNumberOfGuests int                `json:"minNumberOfGuests" bson:"min_number_of_guests"`
	MaxNumberOfGuests int                `json:"maxNumberOfGuests" bson:"max_number_of_guests"`
	Address           Address            `bson:"inline"`
	Images            []Image            `json:"images" bson:"images;"`
	HostID            string             `json:"hostID" bson:"hostID;"`
	Benefits          []string           `json:"benefits" bson:"benefits;"`
}

type Image struct {
	Filename string `json:"filename" bson:"filename"`
	Filepath string `json:"filepath" bson:"filepath"`
}

type Accommodations []*Accommodation

func (u *Accommodations) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Accommodation) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Accommodation) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
