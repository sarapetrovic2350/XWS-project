package model

type Address struct {
	Street       string `bson:"street" json:"street"`
	StreetNumber string `bson:"street_number" json:"street_number"`
	City         string `bson:"city" json:"city"`
	Country      string `bson:"country" json:"country"`
}
