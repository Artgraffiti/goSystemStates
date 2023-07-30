package http

import (
	clientHTTP "GSS/internal/client/http"
	"GSS/internal/metrics"
	"GSS/internal/server/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (server *Server) UploadMetrics(ctx *fiber.Ctx) (err error) {
	var userMetricStorage clientHTTP.UserMetricStorage

	if err := ctx.BodyParser(&userMetricStorage); err != nil {
		return err
	}

	userUUID := userMetricStorage.UUID
	_, ok := storage.GlobalMetricStorage[userUUID]
	if ok {
		timestamp := userMetricStorage.MetricStorage.Timestamp
		for idx, metric := range storage.GlobalMetricStorage[userUUID] {
			if metric.Timestamp == timestamp {
				if len(userMetricStorage.MetricStorage.Float64Data) == 1 {
					for k, v := range userMetricStorage.MetricStorage.Float64Data {
						storage.GlobalMetricStorage[userUUID][idx].Float64Data[k] = v
					}
				} else if len(userMetricStorage.MetricStorage.Uint32Data) == 1 {
					for k, v := range userMetricStorage.MetricStorage.Uint32Data {
						storage.GlobalMetricStorage[userUUID][idx].Uint32Data[k] = v
					}
				} else if len(userMetricStorage.MetricStorage.Uint64Data) == 1 {
					for k, v := range userMetricStorage.MetricStorage.Uint64Data {
						storage.GlobalMetricStorage[userUUID][idx].Uint64Data[k] = v
					}
				}
				return
			}
		}

		storage.GlobalMetricStorage[userUUID] = append(storage.GlobalMetricStorage[userUUID], userMetricStorage.MetricStorage)
	} else {
		storage.GlobalMetricStorage[userUUID] = []metrics.MetricStorage{userMetricStorage.MetricStorage}
	}

	return
}

func (server *Server) MetricsByUUID(ctx *fiber.Ctx) (err error) {
	userUUID, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return
	}

	return ctx.JSON(storage.GlobalMetricStorage[userUUID])
}

func (server *Server) GetUsersMetrics(ctx *fiber.Ctx) (err error) {
	return ctx.JSON(storage.GlobalMetricStorage)
}
