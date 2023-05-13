package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reservation-service/model"
	"reservation-service/service"
)

type KeyProduct struct{}

type ReservationHandler struct {
	logger *log.Logger
	// NoSQL: injecting user service
	reservationService *service.ReservationService
}

// NewUserHandler Injecting the logger makes this code much more testable.
func NewReservationHandler(l *log.Logger, s *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{l, s}
}

func (handler *ReservationHandler) CreateReservation(rw http.ResponseWriter, h *http.Request) {
	fmt.Println("creating")
	var reservation model.Reservation
	err := json.NewDecoder(h.Body).Decode(&reservation)
	if err != nil {
		//TODO log
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(reservation)
	err = handler.reservationService.CreateReservation(&reservation)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (handler *ReservationHandler) GetAllReservations(rw http.ResponseWriter, h *http.Request) {
	reservations, err := handler.reservationService.GetAllReservations()
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if reservations == nil {
		return
	}

	err = reservations.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *ReservationHandler) MiddlewareReservationDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		reservation := &model.Reservation{}
		err := reservation.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, reservation)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (handler *ReservationHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		handler.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
