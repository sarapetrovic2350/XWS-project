package config

import "os"

type Config struct {
	Port                string
	UserDBDomain        string
	UserDBPort          string
	ReservationPort     string
	ReservationDomain   string
	AccommodationDomain string
	AccommodationPort   string
	RatingDomain        string
	RatingPort          string
}

func NewConfig() *Config {
	return &Config{
		Port:                os.Getenv("USER_SERVICE_PORT"),
		UserDBDomain:        os.Getenv("USER_DB_DOMAIN"),
		UserDBPort:          os.Getenv("USER_DB_PORT"),
		ReservationPort:     os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDomain:   os.Getenv("RESERVATION_SERVICE_DOMAIN"),
		AccommodationPort:   os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDomain: os.Getenv("ACCOMMODATION_SERVICE_DOMAIN"),
		RatingPort:          os.Getenv("RATING_SERVICE_PORT"),
		RatingDomain:        os.Getenv("RATING_SERVICE_DOMAIN"),
	}
}
