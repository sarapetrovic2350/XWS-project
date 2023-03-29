package dto

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PurchasedTickets struct {
	Id              string `json:"id"`
	DateOfPurchase  string `json:"dateOfPurchase"`
	DateOfDeparture string `json:"dateOfDeparture"`
	Departure       string `json:"departure"`
	Arrival         string `json:"arrival"`
	NumberOfTickets int    `json:"numberOfTickets"`
	TotalPrice      int    `json:"totalPrice"`
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
