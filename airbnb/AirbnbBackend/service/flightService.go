package service

import (
	"Rest/dto"
	"Rest/model"
	"Rest/repository"
)

type FlightService struct {
	// NoSQL: injecting user repository
	FlightRepo *repository.FlightRepo
}

func NewFlightService(r *repository.FlightRepo) *FlightService {
	return &FlightService{r}
}

func (service *FlightService) CreateFlight(flight *model.Flight) error {
	flight.FlightStatus = model.ONTIME
	err := service.FlightRepo.Insert(flight)
	if err != nil {
		return err
	}
	return nil
}

func (service *FlightService) SearchFlights(searchFlights dto.SearchDTO) model.Flights {
	err := service.FlightRepo.SearchFlights(searchFlights)
	if err != nil {
		return err
	}
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

func (service *FlightService) CancelFlight(id string, flight *model.Flight) error {
	flight.FlightStatus = model.CANCELLED
	err := service.FlightRepo.Update(id, flight)
	if err != nil {
		return err
	}
	return nil
}
