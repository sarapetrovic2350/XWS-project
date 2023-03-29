package service

import (
	"Rest/dto"
	"Rest/model"
	"Rest/repository"
	"time"
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
	var totalPrice = flight.Price * ticket.NumberOfTickets
	ticket.TotalPrice = totalPrice
	ticket.DateOfPurchase = time.Now()
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

func (service *TicketService) GetAllTickets() (model.Tickets, error) {
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

// dobavljanje svih karata koje je kupio jedan korisnik
func (service *TicketService) GetByUserId(userId string) ([]dto.PurchasedTickets, error) {
	//user, _ = service.UserRepo.GetById(userId)
	tickets, err := service.TicketRepo.GetAll()
	var retTickets []dto.PurchasedTickets
	for _, itr := range tickets {
		if itr.IdUser == userId {
			flight, _ := service.FlightRepo.GetById(itr.IdFlight)
			retTickets = append(retTickets, dto.PurchasedTickets{Id: itr.Id.String(), DateOfPurchase: itr.DateOfPurchase.String(),
				DateOfDeparture: flight.DepartureDateTime.String(), Departure: flight.Departure, Arrival: flight.Arrival, NumberOfTickets: itr.NumberOfTickets,
				TotalPrice: itr.TotalPrice})
		}
	}

	if err != nil {
		return nil, err
	}
	return retTickets, nil
}
