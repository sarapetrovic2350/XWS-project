package handler

import (
	"Rest/model"
	"Rest/repository"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type PatientHandler struct {
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *repository.PatientRepo
}

// NewPatientHandler Injecting the logger makes this code much more testable.
func NewPatientHandler(l *log.Logger, r *repository.PatientRepo) *PatientHandler {
	return &PatientHandler{l, r}
}

func (handler *PatientHandler) GetAllPatients(rw http.ResponseWriter, h *http.Request) {
	patients, err := handler.repo.GetAll()
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if patients == nil {
		return
	}

	err = patients.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *PatientHandler) GetPatientById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	patient, err := handler.repo.GetById(id)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if patient == nil {
		http.Error(rw, "Patient with given id not found", http.StatusNotFound)
		handler.logger.Printf("Patient with id: '%s' not found", id)
		return
	}

	err = patient.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *PatientHandler) GetPatientsByName(rw http.ResponseWriter, h *http.Request) {
	name := h.URL.Query().Get("name")

	patients, err := handler.repo.GetByName(name)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if patients == nil {
		return
	}

	err = patients.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *PatientHandler) PostPatient(rw http.ResponseWriter, h *http.Request) {
	patient := h.Context().Value(KeyProduct{}).(*model.Patient)
	handler.repo.Insert(patient)
	rw.WriteHeader(http.StatusCreated)
}

func (handler *PatientHandler) PatchPatient(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	patient := h.Context().Value(KeyProduct{}).(*model.Patient)

	handler.repo.Update(id, patient)
	rw.WriteHeader(http.StatusOK)
}

func (handler *PatientHandler) AddPhoneNumber(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	var phoneNumber string
	d := json.NewDecoder(h.Body)
	d.Decode(&phoneNumber)

	handler.repo.AddPhoneNumber(id, phoneNumber)
	rw.WriteHeader(http.StatusOK)
}

func (handler *PatientHandler) DeletePatient(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	handler.repo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}

func (handler *PatientHandler) AddAnamnesis(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	anamnesis := h.Context().Value(KeyProduct{}).(*model.Anamnesis)

	handler.repo.AddAnamnesis(id, anamnesis)
	rw.WriteHeader(http.StatusOK)
}

func (handler *PatientHandler) AddTherapy(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	therapy := h.Context().Value(KeyProduct{}).(*model.Therapy)

	handler.repo.AddTherapy(id, therapy)
	rw.WriteHeader(http.StatusOK)
}

func (handler *PatientHandler) ChangeAddress(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	address := h.Context().Value(KeyProduct{}).(*model.Address)

	handler.repo.UpdateAddress(id, address)
	rw.WriteHeader(http.StatusOK)
}

func (handler *PatientHandler) ChangePhone(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	index, err := strconv.Atoi(vars["index"])
	if err != nil {
		http.Error(rw, "Unable to decode index", http.StatusBadRequest)
		handler.logger.Fatal(err)
		return
	}

	var phoneNumber string
	d := json.NewDecoder(h.Body)
	d.Decode(&phoneNumber)

	handler.repo.ChangePhone(id, index, phoneNumber)
	rw.WriteHeader(http.StatusOK)
}

func (handler *PatientHandler) Receipt(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	total, err := handler.repo.Receipt(id)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	totalJson := map[string]float64{"total": total}

	e := json.NewEncoder(rw)
	e.Encode(totalJson)
}

func (handler *PatientHandler) Report(rw http.ResponseWriter, h *http.Request) {
	report, err := handler.repo.Report()
	if err != nil {
		handler.logger.Print("Database exception: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(rw)
	e.Encode(report)
}

func (handler *PatientHandler) MiddlewarePatientDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		patient := &model.Patient{}
		err := patient.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, patient)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

// Solution: we added middlewares for Anamnesis, Therapy and Address objects
func (handler *PatientHandler) MiddlewareAnamnesisDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		anamnesis := &model.Anamnesis{}
		err := anamnesis.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, anamnesis)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (handler *PatientHandler) MiddlewareTherapyDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		therapy := &model.Therapy{}
		err := therapy.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, therapy)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (handler *PatientHandler) MiddlewareAddressDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		address := &model.Address{}
		err := address.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, address)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (handler *PatientHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		handler.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
