package main

import (
	"log"

	"GSS/internal/server/config"
	"GSS/internal/server/grpc"
	"GSS/internal/server/http"
)

func main() {
	cfg, err := config.LoadConfig("./config/server/")
	if err != nil {
		log.Fatal(err)
	}

	HTTPServer := http.NewServer(cfg)
	GRPCServer := grpc.NewGRPCServer(cfg)

	GRPCServer.Run()
	HTTPServer.Run()
}
