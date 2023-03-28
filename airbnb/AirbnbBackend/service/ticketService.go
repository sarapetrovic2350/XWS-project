package service

import (
	"Rest/model"
	"Rest/repository"
)

type TicketService struct {
	// NoSQL: injecting user repository
	TicketRepo *repository.TicketRepo
	FlightRepo *repository.FlightRepo
}

func NewTicketService(r *repository.TicketRepo, f *repository.FlightRepo) *TicketService {
	return &TicketService{r, f}
}

func (service *TicketService) CreateTicket(ticket *model.Ticket) (*model.Ticket, error) {
	flight, _ := service.FlightRepo.GetById(ticket.IdFlight)
	var totalPrice = flight.Price * float64(ticket.NumberOfTickets)
	ticket.TotalPrice = totalPrice
	if flight.AvailableSeats >= ticket.NumberOfTickets {
		err := service.TicketRepo.Insert(ticket)
		if err != nil {
			return nil, err
		}
		flight.AvailableSeats = flight.AvailableSeats - ticket.NumberOfTickets
		service.FlightRepo.Update(ticket.IdFlight, flight)
		return ticket, nil
	}
	return nil, nil
}

func (service *TicketService) GetAllFlights() (model.Tickets, error) {
	tickets, err := service.TicketRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

// proveriti za povratni paramaterar
func (service *TicketService) GetById(id string) (*model.Ticket, error) {
	ticket, err := service.TicketRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}
