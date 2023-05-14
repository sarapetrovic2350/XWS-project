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

//func (handler *AccommodationHandler) GetAllAccommodations(rw http.ResponseWriter, h *http.Request) {
//	accommodations, err := handler.accommodationService.GetAllAccommodations()
//	if err != nil {
//	}
//
//	if accommodations == nil {
//		return
//	}
//
//	err = accommodations.ToJSON(rw)
//	if err != nil {
//		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
//		return
//	}
//}
//
//func (handler *AccommodationHandler) GetAccommodationByHostId(rw http.ResponseWriter, h *http.Request) {
//	vars := mux.Vars(h)
//	hostId := vars["id"]
//
//	accommodations, err := handler.accommodationService.GetAccommodationByHostId(hostId)
//
//	if accommodations == nil {
//		return
//	}
//
//	if accommodations == nil {
//		http.Error(rw, "Accommodations with given id not found", http.StatusNotFound)
//		return
//	}
//
//	err = json.NewEncoder(rw).Encode(accommodations)
//	//err = tickets.ToJSON(rw)
//	if err != nil {
//		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
//		return
//	}
//
//	rw.WriteHeader(http.StatusOK)
//	rw.Header().Set("Content-Type", "application/json")
//}
//
//func (handler *AccommodationHandler) SearchAccommodations(rw http.ResponseWriter, h *http.Request) {
//	var dto dto.SearchDTO
//	err := json.NewDecoder(h.Body).Decode(&dto)
//	fmt.Println(dto)
//	accommodations := handler.accommodationService.SearchAccommodation(dto)
//	if accommodations == nil {
//		rw.WriteHeader(http.StatusExpectationFailed)
//		return
//	}
//	err = accommodations.ToJSON(rw)
//	fmt.Println(err)
//	if err != nil {
//		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
//		rw.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	rw.WriteHeader(http.StatusOK)
//	return
//}
//
//func (handler *AccommodationHandler) GetAccommodationById(rw http.ResponseWriter, h *http.Request) {
//	vars := mux.Vars(h)
//	id := vars["id"]
//
//	accommodation, err := handler.accommodationService.GetById(id)
//
//	if accommodation == nil {
//		http.Error(rw, "Accommodation with given id not found", http.StatusNotFound)
//		return
//	}
//	err = accommodation.ToJSON(rw)
//	if err != nil {
//		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
//		return
//	}
//}
