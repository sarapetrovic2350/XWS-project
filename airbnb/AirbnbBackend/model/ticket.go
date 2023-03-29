package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DateOfPurchase  time.Time          `json:"dateOfPurchase" bson:"date_of_purchase"`
	NumberOfTickets int                `json:"numberOfTickets" bson:"number_of_tickets"`
	TotalPrice      int                `json:"totalPrice" bson:"total_price"`
	IdUser          string             `json:"idUser" bson:"id_user"`
	IdFlight        string             `json:"idFlight" bson:"id_flight"`
}

type Tickets []*Ticket

func (u *Tickets) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Ticket) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Ticket) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
