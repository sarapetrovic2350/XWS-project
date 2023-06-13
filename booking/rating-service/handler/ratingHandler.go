package handler

import (
	rating "common/proto/rating-service/pb"
	"context"
	"fmt"
	"rating-service/service"
)

type RatingHandler struct {
	rating.UnimplementedRatingServiceServer
	ratingService *service.RatingService
}

// NewRatingHandler
func NewRatingHandler(s *service.RatingService) *RatingHandler {
	return &RatingHandler{
		ratingService: s,
	}
}
func (handler *RatingHandler) GetAllRatingsHost(ctx context.Context, request *rating.GetAllRatingsHostRequest) (*rating.GetAllRatingsHostResponse, error) {
	fmt.Println("In GetAll grpc api")
	ratingsHost, err := handler.ratingService.GetAllRatingsHost()
	if err != nil {
		return nil, err
	}
	response := &rating.GetAllRatingsHostResponse{
		RatingsHost: []*rating.RatingHost{},
	}
	for _, modelRatingHost := range ratingsHost {
		current := mapRatingHost(modelRatingHost)
		response.RatingsHost = append(response.RatingsHost, current)
	}
	return response, nil
}
func (handler *RatingHandler) CreateRatingForHost(ctx context.Context, request *rating.CreateRatingForHostRequest) (*rating.CreateRatingForHostResponse, error) {
	modelRatingHost := mapNewRatingHost(request.RatingHost)
	err := handler.ratingService.CreateRatingForHost(modelRatingHost)
	if err != nil {
		return nil, err
	}
	return &rating.CreateRatingForHostResponse{
		RatingHost: mapRatingHost(modelRatingHost),
	}, nil
}
