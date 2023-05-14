package main

import (
	"accommodation-service/startup"
	config3 "accommodation-service/startup/config"
)

func main() {
	config := config3.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
