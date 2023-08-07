package http

import (
	"GSS/internal/server/config"
	"context"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HTTPServer struct {
	App    *fiber.App
	Config config.Config
}

func NewHTTPServer(config config.Config) (server *HTTPServer) {
	server = &HTTPServer{
		App:    fiber.New(),
		Config: config,
	}
	server.App.Use(logger.New(logger.Config{
		Format: config.Logger_fmt,
	}))

	server.SetupRoutes()

	return
}

func (server *HTTPServer) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		<-ctx.Done()

		server.App.Shutdown()
		log.Println("HTTP server stoped")

		wg.Done()
	}()

	err := server.App.Listen(server.Config.ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
}
