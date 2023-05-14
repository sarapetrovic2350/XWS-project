package handler

import (
	reservation "common/proto/reservation-service/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reservation-service/model"
	"time"
)

func mapReservation(modelReservation *model.Reservation) *reservation.Reservation {
	reservationPb := &reservation.Reservation{
		Id:                modelReservation.Id.Hex(),
		NumberOfGuests:    int32(modelReservation.NumberOfGuests),
		StartDate:         modelReservation.StartDate.Format("2006-02-01"),
		EndDate:           modelReservation.EndDate.Format("2006-02-01"),
		UserId:            modelReservation.UserId,
		AccommodationId:   modelReservation.AccommodationId,
		ReservationStatus: reservation.ReservationStatus(modelReservation.ReservationStatus),
	}
	return reservationPb
}
func mapNewReservation(reservationPb *reservation.NewReservation) *model.Reservation {
	startDate, _ := time.Parse("2006-01-02", reservationPb.StartDate)
	endDate, _ := time.Parse("2006-01-02", reservationPb.EndDate)
	reservation := &model.Reservation{

		Id:              primitive.NewObjectID(),
		NumberOfGuests:  int(reservationPb.NumberOfGuests),
		StartDate:       startDate,
		EndDate:         endDate,
		UserId:          reservationPb.UserId,
		AccommodationId: reservationPb.AccommodationId,
	}
	return reservation
}
