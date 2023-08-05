package main

import (
	"log"
	"time"

	"GSS/internal/client/config"
	clientGRPC "GSS/internal/client/grpc"
	clientHTTP "GSS/internal/client/http"
	clientMQTT "GSS/internal/client/mqtt"
)

func main() {
	cfg, err := config.LoadConfig("./config/client")
	if err != nil {
		log.Fatal(err)
	}

	/* HTTP Client */

	HTTPUser, err := clientHTTP.NewUser(cfg)
	if err != nil {
		log.Fatal(err)
	}

	HTTPUser.StreamingMetrics()
	time.Sleep(time.Second * 2)

	/* GRPC Client */

	GRPCUser, err := clientGRPC.NewGRPCUser(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer GRPCUser.Close()

	GRPCUser.UploadMetrics()
	time.Sleep(time.Second * 2)

	/* MQTT Client */
	MQTTUser, err := clientMQTT.NewMQTTUser(cfg)
	if err != nil {
		log.Fatal(err)
	}

	MQTTUser.UploadMetrics()
	time.Sleep(time.Second * 2)
}
