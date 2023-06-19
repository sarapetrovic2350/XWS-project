package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type RatingStore interface {
	GetAllRatingsHost() (RatingsHost, error)
	GetAllRatingsAccommodation() (RatingsAccommodation, error)
	GetRatingHostById(id string) (*RatingHost, error)
	GetRatingAccommodationById(id string) (*RatingAccommodation, error)
	GetRatingHost(id primitive.ObjectID) (*RatingHost, error)
	//GetRatingAccommodation(id primitive.ObjectID) (*RatingAccommodation, error)
	InsertRatingHost(rh *RatingHost) error
	InsertRatingAccommodation(ra *RatingAccommodation) error
	//InsertRatingAccommodation(rh *RatingAccommodation) error
	DeleteAll()
	DeleteRatingForHost(id string) error
	DeleteRatingForAccommodation(id string) error
}
