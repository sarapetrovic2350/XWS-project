package main

import (
	"rating-service/startup"
	config2 "rating-service/startup/config"
)

func main() {
	config := config2.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
