package dto

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PurchasedTickets struct {
	Id              string  `json:"id"`
	DateOfPurchase  string  `json:"dateOfPurchase"`
	DateOfDeparture string  `json:"dateOfDeparture"`
	Departure       string  `json:"departure"`
	Arrival         string  `json:"arrival"`
	NumberOfTickets int     `json:"numberOfTickets"`
	TotalPrice      float64 `json:"totalPrice"`
}
