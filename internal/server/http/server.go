package http

import (
	"GSS/internal/server/config"
	"context"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	App    *fiber.App
	Config config.Config
}

func NewServer(config config.Config) (server *Server) {
	server = &Server{
		App:    fiber.New(),
		Config: config,
	}
	server.App.Use(logger.New(logger.Config{
		Format: config.Logger_fmt,
	}))

	server.SetupRoutes()

	return
}

func (server *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		<-ctx.Done()
		server.App.Shutdown()
		wg.Done()
	}()

	err := server.App.Listen(server.Config.ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
}
