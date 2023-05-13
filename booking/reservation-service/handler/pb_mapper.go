package handler

import (
	reservation "common/proto/reservation-service/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reservation-service/model"
)

func mapReservation(modelReservation *model.Reservation) *reservation.Reservation {
	reservationPb := &reservation.Reservation{
		Id:              modelReservation.Id.Hex(),
		NumberOfGuests:  int32(modelReservation.NumberOfGuests),
		StartDate:       timestamppb.New(modelReservation.StartDate),
		EndDate:         timestamppb.New(modelReservation.EndDate),
		UserID:          modelReservation.UserID,
		AccommodationID: modelReservation.AccommodationID,
	}
	return reservationPb
}
