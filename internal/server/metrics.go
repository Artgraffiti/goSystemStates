package server

import (
	"GSS/internal/metrics"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (server *Server) Hi(ctx *fiber.Ctx) (err error) {
	return ctx.SendString("Hi")
}

func (server *Server) addMetrics(ctx *fiber.Ctx) (err error) {
	var data map[uuid.UUID]metrics.Metrics

	err = json.Unmarshal([]byte(ctx.Body()), &data)
	if err != nil {
		return
	}

	for user_id, user_metrics := range data {
		_, ok := server.UsersMetrics[user_id]
		if ok {
			server.UsersMetrics[user_id] = append(server.UsersMetrics[user_id], user_metrics)
		} else {
			server.UsersMetrics[user_id] = []metrics.Metrics{user_metrics}
		}
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
