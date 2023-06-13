package config

import "os"

type Config struct {
	Port                string
	ReservationDBDomain string
	ReservationDBPort   string
	AccommodationPort   string
	AccommodationDomain string
}

func NewConfig() *Config {
	return &Config{
		Port:                os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBDomain: os.Getenv("RESERVATION_DB_DOMAIN"),
		ReservationDBPort:   os.Getenv("RESERVATION_DB_PORT"),
		AccommodationPort:   os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDomain: os.Getenv("ACCOMMODATION_SERVICE_DOMAIN"),
	}
}
