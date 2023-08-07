package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"GSS/internal/server/config"
	"GSS/internal/server/grpc"
	"GSS/internal/server/http"
	"GSS/internal/server/mqtt"
)

func main() {
	ctx, ctxCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer ctxCancel()
	wg := sync.WaitGroup{}

	cfg, err := config.LoadConfig("./config/server/")
	if err != nil {
		log.Fatal(err)
	}

	HTTPServer := http.NewHTTPServer(cfg)
	MQTTServer := mqtt.NewMQTTServer(cfg)
	GRPCServer := grpc.NewGRPCServer(cfg)

	GRPCServer.Run(ctx, &wg)
	MQTTServer.Run(ctx, &wg)
	HTTPServer.Run(ctx, &wg)

	wg.Wait()
}
