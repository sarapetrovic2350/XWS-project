package handler

import (
	"accommodation-service/model"
	"accommodation-service/service"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AvailabilityHandler struct {
	logger *log.Logger
	// NoSQL: injecting user service
	availabilityService *service.AvailabilityService
}

// NewUserHandler Injecting the logger makes this code much more testable.
func NewAvailabilityHandler(l *log.Logger, s *service.AvailabilityService) *AvailabilityHandler {
	return &AvailabilityHandler{l, s}
}

func (handler *AvailabilityHandler) CreateAvailability(rw http.ResponseWriter, h *http.Request) {
	fmt.Println("creating")
	var availability model.Availability
	err := json.NewDecoder(h.Body).Decode(&availability)
	if err != nil {
		//TODO log
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(availability)
	err = handler.availabilityService.CreateAvailability(&availability)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (handler *AvailabilityHandler) GetAllAvailabilities(rw http.ResponseWriter, h *http.Request) {
	availabilities, err := handler.availabilityService.GetAllAvailabilities()
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if availabilities == nil {
		return
	}

	err = availabilities.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *AvailabilityHandler) MiddlewareAccommodationDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		availability := &model.Availability{}
		err := availability.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, availability)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (handler *AvailabilityHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		handler.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
