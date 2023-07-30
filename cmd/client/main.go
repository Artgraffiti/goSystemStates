package main

import (
	"log"

	"GSS/internal/client/config"
	"GSS/internal/client/http"
)

func main() {
	cfg, err := config.LoadConfig("./config/client")
	if err != nil {
		log.Fatal(err)
	}

	user, err := http.NewUser(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = user.StreamingMetrics()
	if err != nil {
		log.Fatal(err)
	}
}
