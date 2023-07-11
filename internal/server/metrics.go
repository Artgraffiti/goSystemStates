package server

import (
	"GSS/internal/metrics"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func (server *Server) Hi(ctx *fiber.Ctx) (err error) {
	return ctx.SendString("Hi")
}

func (server *Server) addMetric(ctx *fiber.Ctx) (err error) {
	var metric metrics.Metric
	err = json.Unmarshal([]byte(ctx.Body()), &metric)
	if err != nil {
		return
	}
	server.Metrics = append(server.Metrics, metric)
	return
}

func (server *Server) getAllMetrics(ctx *fiber.Ctx) (err error) {
	return ctx.JSON(server.Metrics)
}
