package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type RatingStore interface {
	GetAllRatingsHost() (RatingsHost, error)
	//GetAllRatingsAccommodation() (RatingsAccommodation, error)
	GetRatingHost(id primitive.ObjectID) (*RatingHost, error)
	//GetRatingAccommodation(id primitive.ObjectID) (*RatingAccommodation, error)
	InsertRatingHost(rh *RatingHost) error
	//InsertRatingAccommodation(rh *RatingAccommodation) error
	DeleteAll()
	Delete(id string) error
}
