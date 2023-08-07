package main

import (
	"log"
	"time"

	"GSS/internal/client/config"
	"GSS/internal/client/grpc"
	"GSS/internal/client/http"
	"GSS/internal/client/mqtt"
)

func main() {
	cfg, err := config.LoadConfig("./config/client")
	if err != nil {
		log.Fatal(err)
	}

	/* HTTP Client */

	HTTPUser, err := http.NewHTTPUser(cfg)
	if err != nil {
		log.Fatal(err)
	}

	HTTPUser.StreamingMetrics()
	time.Sleep(time.Second * 2)

	/* GRPC Client */

	GRPCUser, err := grpc.NewGRPCUser(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer GRPCUser.Close()

	GRPCUser.UploadMetrics()
	time.Sleep(time.Second * 2)

	/* MQTT Client */

	MQTTUser, err := mqtt.NewMQTTUser(cfg)
	if err != nil {
		log.Fatal(err)
	}

	MQTTUser.UploadMetrics()
	time.Sleep(time.Second * 2)
}
