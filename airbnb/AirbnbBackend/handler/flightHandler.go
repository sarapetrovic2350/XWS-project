package handler

import (
	"Rest/dto"
	"Rest/model"
	"Rest/service"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FlightHandler struct {
	logger        *log.Logger
	flightService *service.FlightService
}

func NewFlightHandler(l *log.Logger, s *service.FlightService) *FlightHandler {
	return &FlightHandler{l, s}
}

func (handler *FlightHandler) CreateFlight(rw http.ResponseWriter, h *http.Request) {
	//var flightDTO dto.FlightDTO
	fmt.Println("ee")
	flight := h.Context().Value(KeyProduct{}).(*dto.FlightDTO)
	fmt.Println(flight)
	if flight == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("prosao prvi if")
	err := handler.flightService.CreateFlight(flight)

	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (handler *FlightHandler) UpdateFlight(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	flight := h.Context().Value(KeyProduct{}).(*model.Flight)

	if err := handler.flightService.UpdateFlight(id, flight); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (handler *FlightHandler) DeleteFlight(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	if err := handler.flightService.DeleteFlight(id); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (handler *FlightHandler) GetAllFlights(rw http.ResponseWriter, h *http.Request) {
	flights, err := handler.flightService.GetAllFlights()
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if flights == nil {
		return
	}

	err = flights.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *FlightHandler) GetFlightById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	flight, err := handler.flightService.GetById(id)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if flight == nil {
		http.Error(rw, "Flight with given id not found", http.StatusNotFound)
		handler.logger.Printf("Flight with id: '%s' not found", id)
		return
	}

	err = flight.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *FlightHandler) MiddlewareUserDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		flight := &model.Flight{}
		err := flight.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, flight)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (handler *FlightHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		handler.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
