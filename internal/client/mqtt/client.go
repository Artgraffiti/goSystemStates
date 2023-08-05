package mqtt

import (
	"GSS/internal/client/config"
	"log"

	"github.com/google/uuid"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTUser struct {
	UUID   uuid.UUID
	Config config.Config
	client mqtt.Client
}

func NewMQTTUser(config config.Config) (user *MQTTUser, err error) {
	userUUID, err := uuid.Parse(config.UUID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Init MQTTUser(uuid): %s", userUUID)

	address := "tcp://" + config.MQTTServerAddr
	opts := mqtt.NewClientOptions().
		AddBroker(address).
		SetClientID(config.UUID)

	user = &MQTTUser{
		UUID:   userUUID,
		Config: config,
		client: mqtt.NewClient(opts),
	}

	if token := user.client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return
}
