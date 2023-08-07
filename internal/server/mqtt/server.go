package mqtt

import (
	"GSS/internal/server/config"
	"context"
	"log"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTServer struct {
	Config config.Config
	client mqtt.Client
}

func NewMQTTServer(config config.Config) (server *MQTTServer) {
	address := "tcp://" + config.MQTTServerAddr
	opts := mqtt.NewClientOptions().
		AddBroker(address).
		SetClientID("myServer(client)")

	server = &MQTTServer{
		Config: config,
		client: mqtt.NewClient(opts),
	}
	return
}

func (server *MQTTServer) Run(ctx context.Context, wg *sync.WaitGroup) {
	log.Println("MQTT message server(receiver) started")
	wg.Add(1)

	go func() {
		<-ctx.Done()

		server.client.Disconnect(250)
		log.Println("MQTT server stoped")

		wg.Done()
	}()

	if token := server.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	server.SetupRoutes()
}
