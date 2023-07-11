package client

import (
	"GSS/internal/client/config"
	"GSS/internal/metrics"
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	uuid   uint64
	Client *fiber.Client
	Config config.Config
}

func NewUser(config config.Config) (user *User) {
	user = &User{
		uuid:   1,
		Client: fiber.AcquireClient(),
		Config: config,
	}
	return
}

func (user *User) SendMetrics() (err error) {
	agent := user.Client.Post("http://" + user.Config.ServerAddr)

	metric, err := metrics.Get()
	if err != nil {
		return
	}

	err = agent.JSON(metric).Parse()
	if err != nil {
		return
	}

	_, _, errs := agent.String()
	if errs != nil {
		return errs[0]
	}
	log.Printf("Sending metric by user_id: %d\n", user.uuid)
	return
}
