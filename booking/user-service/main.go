package main

import (
	"user-service/startup"
	config2 "user-service/startup/config"
)

func main() {
	config := config2.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
