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
	modelReservation := mapNewReservation(request.Reservation)
	err := handler.reservationService.CreateReservation(modelReservation)
	if err != nil {
		return nil, err
	}
	return &reservation.CreateReservationResponse{
		Reservation: mapReservation(modelReservation),
	}, nil
}
func (handler *ReservationHandler) GetReservationsByAccommodationId(ctx context.Context, request *reservation.GetReservationsByAccommodationRequest) (*reservation.GetReservationsByAccommodationResponse, error) {
	fmt.Println("In GetReservationsByAccommodationId grpc api")
	fmt.Println(request)
	activeReservations, err := handler.reservationService.GetReservationsByAccommodationId(request.Id)
	if err != nil {
		return nil, err
	}
	response := &reservation.GetReservationsByAccommodationResponse{
		Reservations: []*reservation.Reservation{},
	}
	for _, modelReservation := range activeReservations {
		current := mapReservation(modelReservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) GetActiveReservationsByGuestId(ctx context.Context, request *reservation.GetActiveReservationsRequest) (*reservation.GetActiveReservationsResponse, error) {
	fmt.Println("In GetActiveReservationsByGuestId grpc api")
	fmt.Println(request)
	activeReservations, err := handler.reservationService.GetActiveReservationsByGuestId(request.Id)
	if err != nil {
		return nil, err
	}
	response := &reservation.GetActiveReservationsResponse{
		Reservations: []*reservation.Reservation{},
	}
	for _, modelReservation := range activeReservations {
		current := mapReservation(modelReservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) GetReservationsByUserId(ctx context.Context, request *reservation.GetReservationsByUserIdRequest) (*reservation.GetReservationsByUserIdResponse, error) {
	fmt.Println("In GetReservationsByUserId grpc api")
	fmt.Println(request)
	reservations, err := handler.reservationService.GetReservationsByUserId(request.Id)
	if err != nil {
		return nil, err
	}
	response := &reservation.GetReservationsByUserIdResponse{
		Reservations: []*reservation.Reservation{},
	}
	for _, modelReservation := range reservations {
		current := mapReservation(modelReservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) GetActiveReservationsByHostId(ctx context.Context, request *reservation.GetActiveReservationsRequest) (*reservation.GetActiveReservationsResponse, error) {
	fmt.Println("In GetActiveReservationsByHostId grpc api")
	fmt.Println(request)
	activeReservations, err := handler.reservationService.GetActiveReservationsByHostId(request.Id)
	if err != nil {
		return nil, err
	}
	response := &reservation.GetActiveReservationsResponse{
		Reservations: []*reservation.Reservation{},
	}
	for _, modelReservation := range activeReservations {
		current := mapReservation(modelReservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (handler *ReservationHandler) DeletePendingReservationByGuest(ctx context.Context, request *reservation.DeleteReservationRequest) (*reservation.DeleteReservationResponse, error) {
	deletedReservation, err := handler.reservationService.GetById(request.Id)
	err = handler.reservationService.Delete(request.Id)
	if err != nil {
		return nil, err
	}
	return &reservation.DeleteReservationResponse{
		Reservation: mapReservation(deletedReservation)}, nil
}
func (handler *ReservationHandler) CancelReservationByGuest(ctx context.Context, request *reservation.CancelReservationRequest) (*reservation.CancelReservationResponse, error) {
	canceledReservation, err := handler.reservationService.CancelReservation(request.Id)
	if err != nil {
		return nil, err
	}
	return &reservation.CancelReservationResponse{
		Reservation: mapReservation(canceledReservation)}, nil
}
func (handler *ReservationHandler) GetPendingReservationsForHost(ctx context.Context, request *reservation.GetPendingReservationsForHostRequest) (*reservation.GetPendingReservationsForHostResponse, error) {
	fmt.Println("In GetPendingReservationsForHost grpc api")
	fmt.Println(request)
	pendingReservations, err := handler.reservationService.GetPendingReservationsForHost(request.Id)
	if err != nil {
		return nil, err
	}
	response := &reservation.GetPendingReservationsForHostResponse{
		Reservations: []*reservation.Reservation{},
	}
	for _, modelReservation := range pendingReservations {
		current := mapReservation(modelReservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
func (handler *ReservationHandler) GetReservationsForHost(ctx context.Context, request *reservation.GetReservationsForHostRequest) (*reservation.GetReservationsForHostResponse, error) {
	fmt.Println("In GetReservationsForHost grpc api")
	fmt.Println(request)
	reservations, err := handler.reservationService.GetReservationsForHost(request.Id)
	if err != nil {
		return nil, err
	}
	response := &reservation.GetReservationsForHostResponse{
		Reservations: []*reservation.Reservation{},
	}
	for _, modelReservation := range reservations {
		current := mapReservation(modelReservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (handler *ReservationHandler) RejectPendingReservationByHost(ctx context.Context, request *reservation.RejectPendingReservationRequest) (*reservation.ReservationResponse, error) {
	rejectedReservation, err := handler.reservationService.RejectReservation(request.Id)
	if err != nil {
		return nil, err
	}
	return &reservation.ReservationResponse{
		Reservation: mapReservation(rejectedReservation)}, nil
}

func (handler *ReservationHandler) AcceptPendingReservationByHost(ctx context.Context, request *reservation.AcceptPendingReservationRequest) (*reservation.ReservationResponse, error) {
	acceptedReservation, err := handler.reservationService.AcceptReservation(request.Id)
	if err != nil {
		return nil, err
	}
	return &reservation.ReservationResponse{
		Reservation: mapReservation(acceptedReservation)}, nil
}

func (handler *ReservationHandler) GetNumberOfPastReservationsByHostId(ctx context.Context, request *reservation.GetNumberOfPastReservationsByHostRequest) (*reservation.GetNumberOfPastReservationsByHostResponse, error) {
	fmt.Println("In GetAll grpc api")
	numberOfReservations, err := handler.reservationService.GetNumberOfPastReservationsByHostId(request.Id)
	if err != nil {
		return nil, err
	}
	response := &reservation.GetNumberOfPastReservationsByHostResponse{
		NumReservations: numberOfReservations,
	}

	return response, nil
}
