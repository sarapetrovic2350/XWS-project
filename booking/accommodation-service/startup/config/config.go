package config

import "os"

type Config struct {
	Port                  string
	AccommodationDBDomain string
	AccommodationDBPort   string
}

func NewConfig() *Config {
	return &Config{
		Port:                  os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBDomain: os.Getenv("ACCOMMODATION_DB_DOMAIN"),
		AccommodationDBPort:   os.Getenv("ACCOMMODATION_DB_PORT"),
	}
}
