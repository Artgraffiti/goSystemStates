package main

import (
	"log"

	"GSS/internal/client/config"
	clientGRPC "GSS/internal/client/grpc"
)

func main() {
	cfg, err := config.LoadConfig("./config/client")
	if err != nil {
		log.Fatal(err)
	}

	GRPCUser, err := clientGRPC.NewGRPCUser(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer GRPCUser.Close()

	GRPCUser.UploadMetrics()
}
