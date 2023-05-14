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

//	func (handler *AccommodationHandler) CreateAccommodation(rw http.ResponseWriter, h *http.Request) {
//		fmt.Println("creating")
//		var accommodation model.Accommodation
//		err := json.NewDecoder(h.Body).Decode(&accommodation)
//		if err != nil {
//			//TODO log
//			rw.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		fmt.Println(accommodation)
//		err = handler.accommodationService.CreateAccommodation(&accommodation)
//		if err != nil {
//			fmt.Println(err)
//			rw.WriteHeader(http.StatusExpectationFailed)
//		}
//		rw.WriteHeader(http.StatusCreated)
//		rw.Header().Set("Content-Type", "application/json")
//	}
func (handler *AccommodationHandler) CreateAccommodation(ctx context.Context, request *accommodation.CreateAccommodationRequest) (*accommodation.CreateAccommodationResponse, error) {
	fmt.Println("In CreateAccommodation grpc api")
	fmt.Print("Request.Accommodation: ")
	fmt.Println(request.Accommodation)
	modelAccommodation := mapNewAccommodation(request.Accommodation)
	fmt.Print("user after mapping: ")
	fmt.Println(modelAccommodation)
	err := handler.accommodationService.CreateAccommodation(modelAccommodation)
	if err != nil {
		return nil, err
	}
	return &accommodation.CreateAccommodationResponse{
		Accommodation: mapAccommodation(modelAccommodation),
	}, nil
}

func (handler *AccommodationHandler) CreateAvailability(ctx context.Context, request *accommodation.CreateAvailabilityRequest) (*accommodation.CreateAvailabilityResponse, error) {
	fmt.Println("In CreateAvailability grpc api")
	fmt.Print("Request.AccommodationId: ")
	fmt.Println(request.AvailabilityForAccommodation.AccommodationId)
	fmt.Println(request.AvailabilityForAccommodation.Availability)
	modelAvailability := mapNewAvailability(request.AvailabilityForAccommodation.Availability)
	fmt.Print("availability after mapping: ")
	fmt.Println(modelAvailability)
	err := handler.accommodationService.AddAvailabilityForAccommodation(request)
	if err != nil {
		return nil, err
	}
	return &accommodation.CreateAvailabilityResponse{
		Availability: mapAvailabilityPb(modelAvailability),
	}, nil
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
