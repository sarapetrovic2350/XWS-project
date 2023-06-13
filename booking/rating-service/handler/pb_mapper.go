package handler

import (
	rating "common/proto/rating-service/pb"
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
