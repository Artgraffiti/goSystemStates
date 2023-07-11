package server

import (
	"GSS/internal/metrics"
	"GSS/internal/server/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	App     *fiber.App
	Config  config.Config
	Metrics []metrics.Metric
}

func NewServer(config config.Config) (server *Server) {
	server = &Server{
		App:     fiber.New(),
		Config:  config,
		Metrics: make([]metrics.Metric, 0),
	}
	server.App.Use(logger.New(logger.Config{
		Format: config.Logger_fmt,
	}))

	server.SetupRoutes()

	return
}

func (server *Server) Run() {
	err := server.App.Listen(server.Config.ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
}
