package config

import "os"

type Config struct {
	Port                     string
	ReservationDBDomain      string
	ReservationDBPort        string
	AccommodationPort        string
	AccommodationDomain      string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                     os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBDomain:      os.Getenv("RESERVATION_DB_DOMAIN"),
		ReservationDBPort:        os.Getenv("RESERVATION_DB_PORT"),
		AccommodationPort:        os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDomain:      os.Getenv("ACCOMMODATION_SERVICE_DOMAIN"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		DeleteUserCommandSubject: os.Getenv("DELETE_USER_COMMAND_SUBJECT"),
		DeleteUserReplySubject:   os.Getenv("DELETE_USER_REPLY_SUBJECT"),
	}
}
