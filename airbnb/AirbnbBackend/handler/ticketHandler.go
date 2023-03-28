package handler

import (
	"Rest/model"
	"Rest/service"
	"context"
	"log"
	"net/http"
)

type TicketHandler struct {
	logger        *log.Logger
	ticketService *service.TicketService
}

func NewTicketHandler(l *log.Logger, s *service.TicketService) *TicketHandler {
	return &TicketHandler{l, s}
}

func (handler *TicketHandler) CreateTicket(rw http.ResponseWriter, h *http.Request) {
	ticket := h.Context().Value(KeyProduct{}).(*model.Ticket)
	createdTicket, err := handler.ticketService.CreateTicket(ticket)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if createdTicket == nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (handler *TicketHandler) GetAllTickets(rw http.ResponseWriter, h *http.Request) {
	tickets, err := handler.ticketService.GetAllFlights()
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if tickets == nil {
		return
	}

	err = tickets.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *TicketHandler) MiddlewareTicketDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		ticket := &model.Ticket{}
		err := ticket.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			handler.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, ticket)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (handler *TicketHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		handler.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
