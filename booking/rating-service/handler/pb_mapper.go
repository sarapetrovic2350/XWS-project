package handler

import (
	rating "common/proto/rating-service/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating-service/model"
)

func mapRatingHost(modelRatingHost *model.RatingHost) *rating.RatingHost {
	ratingHostPb := &rating.RatingHost{
		Id:      modelRatingHost.Id.Hex(),
		HostId:  modelRatingHost.HostId,
		GuestId: modelRatingHost.GuestId,
		Rate:    int32(modelRatingHost.Rate),
	}
	return ratingHostPb
}
func mapNewRatingHost(ratingHostPb *rating.NewRatingHost) *model.RatingHost {
	ratingHost := &model.RatingHost{

		Id:      primitive.NewObjectID(),
		HostId:  ratingHostPb.HostId,
		GuestId: ratingHostPb.GuestId,
		Rate:    uint32(ratingHostPb.Rate),
	}
	return ratingHost
}
func mapUpdatedRatingHost(ratingHostPb *rating.RatingHost) *model.RatingHost {
	RatingHostId, _ := primitive.ObjectIDFromHex(ratingHostPb.Id)
	ratingHost := &model.RatingHost{
		Id:      RatingHostId,
		HostId:  ratingHostPb.HostId,
		GuestId: ratingHostPb.GuestId,
		Rate:    uint32(ratingHostPb.Rate),
	}
	return ratingHost
}

func mapNewRatingAccommodation(ratingAccommodationPb *rating.NewRatingAccommodation) *model.RatingAccommodation {
	ratingAccommodation := &model.RatingAccommodation{

		Id:              primitive.NewObjectID(),
		AccommodationId: ratingAccommodationPb.AccommodationId,
		GuestId:         ratingAccommodationPb.GuestId,
		Rate:            uint32(ratingAccommodationPb.Rate),
	}
	return ratingAccommodation
}
func mapRatingAccommodation(modelRatingAccommodation *model.RatingAccommodation) *rating.RatingAccommodation {
	ratingAccommodationPb := &rating.RatingAccommodation{
		Id:              modelRatingAccommodation.Id.Hex(),
		AccommodationId: modelRatingAccommodation.AccommodationId,
		GuestId:         modelRatingAccommodation.GuestId,
		Rate:            int32(modelRatingAccommodation.Rate),
	}
	return ratingAccommodationPb
}
