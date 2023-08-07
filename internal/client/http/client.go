package http

import (
	"GSS/internal/client/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HTTPUser struct {
	UUID   uuid.UUID
	Client *fiber.Client
	Config config.Config
}

func NewHTTPUser(config config.Config) (user *HTTPUser, err error) {
	userUUID, err := uuid.Parse(config.UUID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Init user(uuid): %s", userUUID)
	user = &HTTPUser{
		UUID:   userUUID,
		Client: fiber.AcquireClient(),
		Config: config,
	}
	return
}
