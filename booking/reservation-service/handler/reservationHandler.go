package handler

import (
	reservation "common/proto/reservation-service/pb"
	"context"
	"fmt"
	"reservation-service/service"
)

type KeyProduct struct{}

type ReservationHandler struct {
	reservation.UnimplementedReservationServiceServer
	reservationService *service.ReservationService
}

// NewReservationHandler Injecting the logger makes this code much more testable.
func NewReservationHandler(s *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		reservationService: s,
	}
}
func (handler *ReservationHandler) GetAll(ctx context.Context, request *reservation.GetAllRequest) (*reservation.GetAllResponse, error) {
	fmt.Println("In GetAll grpc api")
	reservations, err := handler.reservationService.GetAllReservations()
	if err != nil {
		return nil, err
	}
	response := &reservation.GetAllResponse{
		Reservations: []*reservation.Reservation{},
	}
	for _, modelReservation := range reservations {
		current := mapReservation(modelReservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) CreateReservation(ctx context.Context, request *reservation.CreateReservationRequest) (*reservation.CreateReservationResponse, error) {
	fmt.Println("In CreateReservation grpc api")
	fmt.Print("Request.Reservation: ")
	fmt.Println(request.Reservation)
	modelReservation := mapNewReservation(request.Reservation)
	fmt.Print("reservation after mapping: ")
	fmt.Println(modelReservation)
	err := handler.reservationService.CreateReservation(modelReservation)
	if err != nil {
		return nil, err
	}
	return &reservation.CreateReservationResponse{
		Reservation: mapReservation(modelReservation),
	}, nil
}
func (handler *ReservationHandler) GetReservationsByUserId(ctx context.Context, request *reservation.GetUserReservationsRequest) (*reservation.GetUserReservationsResponse, error) {
	fmt.Println("In GetReservationsByUserId grpc api")
	fmt.Println(request)
	reservations, err := handler.reservationService.GetReservationsByUserId(request.Id)
	if err != nil {
		return nil, err
	}
	response := &reservation.GetUserReservationsResponse{
		Reservations: []*reservation.Reservation{},
	}
	for _, modelReservation := range reservations {
		current := mapReservation(modelReservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

//func (handler *ReservationHandler) CreateReservation(rw http.ResponseWriter, h *http.Request) {
//	fmt.Println("creating")
//	var reservation model.Reservation
//	err := json.NewDecoder(h.Body).Decode(&reservation)
//	if err != nil {
//		//TODO log
//		rw.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	fmt.Println(reservation)
//	err = handler.reservationService.CreateReservation(&reservation)
//	if err != nil {
//		fmt.Println(err)
//		rw.WriteHeader(http.StatusExpectationFailed)
//	}
//	rw.WriteHeader(http.StatusCreated)
//	rw.Header().Set("Content-Type", "application/json")
//}

//func (handler *ReservationHandler) GetAllReservations(rw http.ResponseWriter, h *http.Request) {
//	reservations, err := handler.reservationService.GetAllReservations()
//	if reservations == nil {
//		return
//	}
//
//	err = reservations.ToJSON(rw)
//	if err != nil {
//		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
//		return
//	}
//}

//func (handler *ReservationHandler) DeleteReservation(ctx context.Context, request *reservation.) (*user.DeleteUserResponse, error) {
//	formatedId, err := primitive.ObjectIDFromHex(request.Id)
//	if err != nil {
//		return nil, err
//	}
//	deletedUser, err := handler.reservationService.Get(formatedId)
//	err = handler.reservationService.Delete(request)
//	if err != nil {
//		return nil, err
//	}
//	return &user.DeleteUserResponse{
//		User: mapUser(deletedUser)}, nil
//}
