package config

import "os"

type Config struct {
	Port              string
	UserDBDomain      string
	UserDBPort        string
	ReservationPort   string
	ReservationDomain string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("USER_SERVICE_PORT"),
		UserDBDomain:      os.Getenv("USER_DB_DOMAIN"),
		UserDBPort:        os.Getenv("USER_DB_PORT"),
		ReservationPort:   os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDomain: os.Getenv("RESERVATION_SERVICE_DOMAIN"),
	}
}