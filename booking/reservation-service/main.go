package main

import (
	"reservation-service/startup"
	config2 "reservation-service/startup/config"
)

func main() {
	config := config2.NewConfig()
	server := startup.NewServer(config)
	server.Start()

}
