package client

import (
	"GSS/internal/client/config"
	"GSS/internal/metrics"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	uuid   uuid.UUID
	Client *fiber.Client
	Config config.Config
}

func NewUser(config config.Config) (user *User, err error) {
	user_uuid, err := uuid.Parse(config.UUID)
	log.Printf("Initialize user(uuid): %s", user_uuid)
	user = &User{
		uuid:   user_uuid,
		Client: fiber.AcquireClient(),
		Config: config,
	}
	return
}

func (user *User) SendMetrics() (err error) {
	agent := user.Client.Post("http://" + user.Config.ServerAddr)

	userMetrics, err := metrics.Get()
	if err != nil {
		return
	}

	request := map[uuid.UUID]metrics.Metrics{user.uuid: userMetrics}
	err = agent.JSON(request).Parse()
	if err != nil {
		return
	}

	statusCode, _, errs := agent.String()
	if errs != nil {
		return errs[0]
	}
	log.Printf("Sending metric by user_id: %s\n With status code: %d", user.uuid.String(), statusCode)
	return
}
