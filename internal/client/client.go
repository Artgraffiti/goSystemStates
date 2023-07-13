package client

import (
	"GSS/internal/client/config"
	"GSS/internal/metrics"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserMetricMap struct {
	UUID      uuid.UUID         `json:"uuid"`
	MetricMap metrics.MetricMap `json:"metrics"`
}

type User struct {
	uuid.UUID
	Client *fiber.Client
	Config config.Config
}

func NewUser(config config.Config) (user *User, err error) {
	user_uuid, err := uuid.Parse(config.UUID)
	log.Printf("Init user(uuid): %s", user_uuid)
	user = &User{
		UUID:   user_uuid,
		Client: fiber.AcquireClient(),
		Config: config,
	}
	return
}

func (user *User) SendMetricMap() (err error) {
	agent := user.Client.Post("http://" + user.Config.ServerAddr)

	mMap, err := metrics.Get()
	if err != nil {
		return
	}

	request := UserMetricMap{
		UUID:      user.UUID,
		MetricMap: mMap,
	}
	err = agent.JSON(request).Parse()
	if err != nil {
		return
	}

	statusCode, _, errs := agent.String()
	if errs != nil {
		return errs[0]
	}
	log.Printf("Sending metric...\nStatus code: %d", statusCode)
	return
}
