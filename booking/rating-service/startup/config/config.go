package config

import "os"

type Config struct {
	Port                string
	RatingDBDomain      string
	RatingDBPort        string
	ReservationPort     string
	ReservationDomain   string
	AccommodationDomain string
	AccommodationPort   string
	UserDomain          string
	UserPort            string
}

func NewConfig() *Config {
	return &Config{
		Port:                os.Getenv("RATING_SERVICE_PORT"),
		RatingDBDomain:      os.Getenv("RATING_DB_DOMAIN"),
		RatingDBPort:        os.Getenv("RATING_DB_PORT"),
		ReservationPort:     os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDomain:   os.Getenv("RESERVATION_SERVICE_DOMAIN"),
		AccommodationPort:   os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDomain: os.Getenv("ACCOMMODATION_SERVICE_DOMAIN"),
		UserPort:            os.Getenv("USER_SERVICE_PORT"),
		UserDomain:          os.Getenv("USER_SERVICE_DOMAIN"),
	}
}
