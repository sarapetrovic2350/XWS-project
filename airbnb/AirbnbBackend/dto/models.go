package dto

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FlightDTO struct {
	Id                 string `json:"id"`
	DepartureDate      string `json:"departureDate" `
	DepartureTime      string `json:"departureTime" `
	ArrivalDate        string `json:"arrivalDate `
	ArrivalTime        string `json:"arrivalTime" `
	Departure          string `json:"departure" `
	Arrival            string `json:"arrival" `
	Price              int    `json:"price" `
	TotalNumberOfSeats int    `json:"totalNumberOfSeats" `
	AvailableSeats     int    `json:"availableSeats" `
}
