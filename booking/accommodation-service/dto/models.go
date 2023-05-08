package dto

type SearchDTO struct {
	Country string `json:"country"`
	City    string `json:"city"`
	//Date           time.Time `json:"date"`
	NumberOfGuests int `json:"numberOfGuests"`
}
