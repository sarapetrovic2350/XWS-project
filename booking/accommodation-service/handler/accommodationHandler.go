package handler

import (
	"accommodation-service/service"
	accommodation "common/proto/accommodation-service/pb"
	"context"
	"fmt"
)

type AccommodationHandler struct {
	accommodation.UnimplementedAccommodationServiceServer
	accommodationService *service.AccommodationService
}

// NewAccommodationHandler Injecting the logger makes this code much more testable.
func NewAccommodationHandler(s *service.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		accommodationService: s,
	}
}
func (handler *AccommodationHandler) GetAll(ctx context.Context, request *accommodation.GetAllRequest) (*accommodation.GetAllResponse, error) {
	fmt.Println("In GetAll grpc api")
	accommodations, err := handler.accommodationService.GetAllAccommodations()
	if err != nil {
		return nil, err
	}
	response := &accommodation.GetAllResponse{
		Accommodations: []*accommodation.Accommodation{},
	}
	for _, modelAccommodation := range accommodations {
		current := mapAccommodation(modelAccommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil
}
func (handler *AccommodationHandler) GetAccommodationsByHostId(ctx context.Context, request *accommodation.GetAccommodationsByHostIdRequest) (*accommodation.GetAccommodationsByHostIdResponse, error) {
	fmt.Println("In GetAccommodationsByHostId grpc api")
	accommodations, err := handler.accommodationService.GetAccommodationByHostId(request.HostId)
	if err != nil {
		return nil, err
	}
	response := &accommodation.GetAccommodationsByHostIdResponse{
		Accommodations: []*accommodation.Accommodation{},
	}
	for _, modelAccommodation := range accommodations {
		current := mapAccommodation(modelAccommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil
}

func (handler *AccommodationHandler) CreateAccommodation(ctx context.Context, request *accommodation.CreateAccommodationRequest) (*accommodation.CreateAccommodationResponse, error) {
	fmt.Println("In CreateAccommodation grpc api")
	fmt.Print("Request.Accommodation: ")
	fmt.Println(request.Accommodation)
	modelAccommodation := mapNewAccommodation(request.Accommodation)
	fmt.Print("accommodation after mapping: ")
	fmt.Println(modelAccommodation)
	err := handler.accommodationService.CreateAccommodation(modelAccommodation)
	if err != nil {
		return nil, err
	}
	return &accommodation.CreateAccommodationResponse{
		Accommodation: mapAccommodation(modelAccommodation),
	}, nil
}

func (handler *AccommodationHandler) CreateAvailability(ctx context.Context, request *accommodation.CreateAvailabilityRequest) (*accommodation.CreateAccommodationResponse, error) {
	fmt.Println("In CreateAvailability grpc api")
	fmt.Print("Request.AccommodationId: ")
	fmt.Println(request.AvailabilityForAccommodation.AccommodationId)
	fmt.Println(request.AvailabilityForAccommodation.Availability)
	accommodationForUpdate, err := handler.accommodationService.GetById(request.AvailabilityForAccommodation.AccommodationId)
	fmt.Println(accommodationForUpdate.Availabilities)
	if err != nil {
		return nil, err
	}

	err = handler.accommodationService.AddAvailabilityForAccommodation(accommodationForUpdate, mapNewAvailability(request.AvailabilityForAccommodation.Availability))
	if err != nil {
		return nil, err
	}

	updated, err := handler.accommodationService.GetById(request.AvailabilityForAccommodation.AccommodationId)
	if err != nil {
		return nil, err
	}

	response := &accommodation.CreateAccommodationResponse{
		Accommodation: mapAccommodation(updated),
	}

	return response, nil
}

func (handler AccommodationHandler) Search(ctx context.Context, request *accommodation.GetAccommodationsByParamsRequest) (*accommodation.GetAccommodationsByParamsResponse, error) {

	accommodations := handler.accommodationService.SearchAccommodation(request)
	if accommodations == nil {
		return nil, nil
	}

	response := &accommodation.GetAccommodationsByParamsResponse{
		Accommodations: []*accommodation.Accommodation{},
	}

	for _, modelAccommodation := range accommodations {
		current := mapAccommodation(modelAccommodation)
		response.Accommodations = append(response.Accommodations, current)
	}

	return response, nil
}
