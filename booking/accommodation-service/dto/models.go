package dto

import "time"

type SearchDTO struct {
	Country        string    `json:"country"`
	City           string    `json:"city"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
	NumberOfGuests int       `json:"numberOfGuests"`
}
