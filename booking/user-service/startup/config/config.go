package config

import "os"

type Config struct {
	Port                     string
	UserDBDomain             string
	UserDBPort               string
	ReservationPort          string
	ReservationDomain        string
	AccommodationDomain      string
	AccommodationPort        string
	RatingDomain             string
	RatingPort               string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                     os.Getenv("USER_SERVICE_PORT"),
		UserDBDomain:             os.Getenv("USER_DB_DOMAIN"),
		UserDBPort:               os.Getenv("USER_DB_PORT"),
		ReservationPort:          os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDomain:        os.Getenv("RESERVATION_SERVICE_DOMAIN"),
		AccommodationPort:        os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDomain:      os.Getenv("ACCOMMODATION_SERVICE_DOMAIN"),
		RatingPort:               os.Getenv("RATING_SERVICE_PORT"),
		RatingDomain:             os.Getenv("RATING_SERVICE_DOMAIN"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		DeleteUserCommandSubject: os.Getenv("DELETE_USER_COMMAND_SUBJECT"),
		DeleteUserReplySubject:   os.Getenv("DELETE_USER_REPLY_SUBJECT"),
	}
}
