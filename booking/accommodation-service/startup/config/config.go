package config

import "os"

type Config struct {
	Port                     string
	AccommodationDBDomain    string
	AccommodationDBPort      string
	ReservationDomain        string
	ReservationPort          string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                     os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBDomain:    os.Getenv("ACCOMMODATION_DB_DOMAIN"),
		AccommodationDBPort:      os.Getenv("ACCOMMODATION_DB_PORT"),
		ReservationDomain:        os.Getenv("RESERVATION_SERVICE_DOMAIN"),
		ReservationPort:          os.Getenv("RESERVATION_SERVICE_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		DeleteUserCommandSubject: os.Getenv("DELETE_USER_COMMAND_SUBJECT"),
		DeleteUserReplySubject:   os.Getenv("DELETE_USER_REPLY_SUBJECT"),
	}
}
