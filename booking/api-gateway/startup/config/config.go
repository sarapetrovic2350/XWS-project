package config

import "os"

type Config struct {
	Port              string
	UserDomain        string
	UserPort          string
	ReservationDomain string
	ReservationPort   string
	AllowedCorsOrigin string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("GATEWAY_PORT"),
		UserPort:          os.Getenv("USER_SERVICE_PORT"),
		UserDomain:        os.Getenv("USER_SERVICE_DOMAIN"),
		ReservationDomain: os.Getenv("RESERVATION_SERVICE_DOMAIN"),
		ReservationPort:   os.Getenv("RESERVATION_SERVICE_PORT"),
		AllowedCorsOrigin: "http://localhost:4200",
	}
}
