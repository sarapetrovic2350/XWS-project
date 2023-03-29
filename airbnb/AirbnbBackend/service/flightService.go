package service

import (
	"Rest/dto"
	"Rest/model"
	"Rest/repository"
	"fmt"
	"time"
)

type FlightService struct {
	// NoSQL: injecting user repository
	FlightRepo *repository.FlightRepo
}

func NewFlightService(r *repository.FlightRepo) *FlightService {
	return &FlightService{r}
}

func (service *FlightService) CreateFlight(flightDTO *dto.FlightDTO) error {
	println(flightDTO)
	d1, _ := time.Parse("2006-01-02", flightDTO.DepartureDate)
	d2, _ := time.Parse("2006-01-02", flightDTO.ArrivalDate)

	flight := model.Flight{DepartureTime1: d1, ArrivalTime1: d2,
		Departure: flightDTO.Departure, Arrival: flightDTO.Arrival, Price: flightDTO.Price, TotalNumberOfSeats: flightDTO.TotalNumberOfSeats,
		AvailableSeats: flightDTO.AvailableSeats}
	fmt.Println(flight)
	fmt.Println("kreiran let")

	service.FlightRepo.Insert(&flight)
	return nil
}

func (service *FlightService) GetAllFlights() (model.Flights, error) {
	flights, err := service.FlightRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return flights, nil
}

// proveriti za povratni paramaterar
func (service *FlightService) GetById(id string) (*model.Flight, error) {
	flight, err := service.FlightRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return flight, nil
}

func (service *FlightService) UpdateFlight(id string, flight *model.Flight) error {
	err := service.FlightRepo.Update(id, flight)
	if err != nil {
		return err
	}
	return nil
}

func (service *FlightService) DeleteFlight(id string) error {
	err := service.FlightRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
