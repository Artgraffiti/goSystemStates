package main

import (
	"log"

	"GSS/internal/server"
	"GSS/internal/server/config"
)

func main() {
	cfg, err := config.LoadConfig("./config/server/")
	if err != nil {
		log.Fatal(err)
	}

	HTTPServer := server.NewServer(cfg)
	GRPCServer := server.NewGRPCServer(cfg)

	GRPCServer.Run()
	HTTPServer.Run()
}
