package config

import "os"

type Config struct {
	Port                  string
	AccommodationDBDomain string
	AccommodationDBPort   string
	ReservationDomain     string
	ReservationPort       string
}

func NewConfig() *Config {
	return &Config{
		Port:                  os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBDomain: os.Getenv("ACCOMMODATION_DB_DOMAIN"),
		AccommodationDBPort:   os.Getenv("ACCOMMODATION_DB_PORT"),
		ReservationDomain:     os.Getenv("RESERVATION_SERVICE_DOMAIN"),
		ReservationPort:       os.Getenv("RESERVATION_SERVICE_PORT"),
	}
}
