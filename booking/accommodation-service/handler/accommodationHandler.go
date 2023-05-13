package handler

import (
	"accommodation-service/dto"
	"accommodation-service/model"
	"accommodation-service/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type KeyProduct struct{}

type AccommodationHandler struct {
	logger *log.Logger
	// NoSQL: injecting user service
	accommodationService *service.AccommodationService
}

// NewUserHandler Injecting the logger makes this code much more testable.
func NewAccommodationHandler(l *log.Logger, s *service.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{l, s}
}

func (handler *AccommodationHandler) CreateAccommodation(rw http.ResponseWriter, h *http.Request) {
	fmt.Println("creating")
	var accommodation model.Accommodation
	err := json.NewDecoder(h.Body).Decode(&accommodation)
	if err != nil {
		//TODO log
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(accommodation)
	err = handler.accommodationService.CreateAccommodation(&accommodation)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusExpectationFailed)
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
}

func (handler *AccommodationHandler) GetAllAccommodations(rw http.ResponseWriter, h *http.Request) {
	accommodations, err := handler.accommodationService.GetAllAccommodations()
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if accommodations == nil {
		return
	}

	err = accommodations.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *AccommodationHandler) GetAccommodationByHostId(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	hostId := vars["id"]

	accommodations, err := handler.accommodationService.GetAccommodationByHostId(hostId)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if accommodations == nil {
		return
	}

	if accommodations == nil {
		http.Error(rw, "Accommodations with given id not found", http.StatusNotFound)
		handler.logger.Printf("Accommodations with id: '%s' not found", hostId)
		return
	}

	err = json.NewEncoder(rw).Encode(accommodations)
	//err = tickets.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
}

func (handler *AccommodationHandler) MiddlewareAccommodationDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		accommodation := &model.Accommodation{}
		err := accommodation.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, accommodation)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (handler *AccommodationHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		handler.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func (handler *AccommodationHandler) SearchAccommodations(rw http.ResponseWriter, h *http.Request) {
	var dto dto.SearchDTO
	err := json.NewDecoder(h.Body).Decode(&dto)
	fmt.Println(dto)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}
	accommodations := handler.accommodationService.SearchAccommodation(dto)
	if accommodations == nil {
		rw.WriteHeader(http.StatusExpectationFailed)
		return
	}
	err = accommodations.ToJSON(rw)
	fmt.Println(err)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
	return
}

func (handler *AccommodationHandler) GetAccommodationById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	accommodation, err := handler.accommodationService.GetById(id)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if accommodation == nil {
		http.Error(rw, "Accommodation with given id not found", http.StatusNotFound)
		handler.logger.Printf("Accommodation with id: '%s' not found", id)
		return
	}

	err = accommodation.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}
