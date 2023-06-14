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
	createdRating, err := handler.ratingService.CreateRatingForHost(modelRatingHost)
	if err != nil {
		return nil, err
	}
	return &rating.CreateRatingForHostResponse{
		RatingHost: mapRatingHost(createdRating),
	}, nil
}
func (handler *RatingHandler) CreateRatingForAccommodation(ctx context.Context, request *rating.CreateRatingForAccommodationRequest) (*rating.CreateRatingForAccommodationResponse, error) {
	fmt.Println("CreateRatingForAccommodation handler")
	fmt.Println(request.RatingAccommodation)
	modelRatingAccommodation := mapNewRatingAccommodation(request.RatingAccommodation)
	createdRating, err := handler.ratingService.CreateRatingForAccommodation(modelRatingAccommodation)
	if err != nil {
		return nil, err
	}
	return &rating.CreateRatingForAccommodationResponse{
		RatingAccommodation: mapRatingAccommodation(createdRating),
	}, nil
}
func (handler *RatingHandler) DeleteRatingForHost(ctx context.Context, request *rating.DeleteRatingForHostRequest) (*rating.DeleteRatingForHostResponse, error) {
	deletedRating, err := handler.ratingService.GetRatingHostById(request.Id)
	err = handler.ratingService.DeleteRatingForHost(request.Id)
	if err != nil {
		return nil, err
	}
	return &rating.DeleteRatingForHostResponse{
		RatingHost: mapRatingHost(deletedRating)}, nil
}
func (handler *RatingHandler) UpdateRatingForHost(ctx context.Context, request *rating.UpdateRatingForHostRequest) (*rating.UpdateRatingForHostResponse, error) {
	fmt.Print("Request.RatingHost: ")
	modelRatingHost := mapUpdatedRatingHost(request.RatingHost)
	fmt.Print("rating host after mapping: ")
	fmt.Println(modelRatingHost)
	updatedRating, err := handler.ratingService.UpdateRatingForHost(modelRatingHost)
	if err != nil {
		return nil, err
	}
	return &rating.UpdateRatingForHostResponse{
		RatingHost: mapRatingHost(updatedRating)}, nil
}
