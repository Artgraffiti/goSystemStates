package main

import (
	"log"

	"GSS/internal/client"
	"GSS/internal/client/config"
)

func main() {
	cfg, err := config.LoadConfig("../../config/client")
	if err != nil {
		log.Fatal(err)
	}

	user := client.NewUser(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = user.SendMetrics()
	if err != nil {
		log.Fatal(err)
	}
}
