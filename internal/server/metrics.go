package server

import (
	"GSS/internal/client"
	"GSS/internal/metrics"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (server *Server) Hi(ctx *fiber.Ctx) (err error) {
	return ctx.SendString("Hi")
}

func (server *Server) addMetrics(ctx *fiber.Ctx) (err error) {
	var userMetricMap client.UserMetricMap

	if err := ctx.BodyParser(&userMetricMap); err != nil {
		return err
	}

	_, ok := server.UsersMetrics[userMetricMap.UUID]
	if ok {
		server.UsersMetrics[userMetricMap.UUID] = append(server.UsersMetrics[userMetricMap.UUID], userMetricMap.MetricStorage)
	} else {
		server.UsersMetrics[userMetricMap.UUID] = []metrics.MetricStorage{userMetricMap.MetricStorage}
	}

	return
}

func (server *Server) metricsByUUID(ctx *fiber.Ctx) (err error) {
	user_uuid, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return
	}

	return ctx.JSON(server.UsersMetrics[user_uuid])
}

func (server *Server) getAllMetrics(ctx *fiber.Ctx) (err error) {
	return ctx.JSON(server.UsersMetrics)
}
