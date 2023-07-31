package http

import (
	clientHTTP "GSS/internal/client/http"
	"GSS/internal/metrics"
	"GSS/internal/server/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func unloadUserMetricStorageToGlobalMetricStorage(data clientHTTP.UserMetricStorage) (err error) {
	userUUID := data.UUID
	userMetricStorage, ok := storage.GlobalMetricStorage[userUUID]
	if ok {
		timestamp := data.MetricStorage.Timestamp
		for idx, metric := range userMetricStorage {
			if metric.Timestamp == timestamp {
				if len(data.MetricStorage.Float64Data) == 1 {
					for k, v := range data.MetricStorage.Float64Data {
						storage.GlobalMetricStorage[userUUID][idx].Float64Data[k] = v
					}
				} else if len(data.MetricStorage.Uint32Data) == 1 {
					for k, v := range data.MetricStorage.Uint32Data {
						storage.GlobalMetricStorage[userUUID][idx].Uint32Data[k] = v
					}
				} else if len(data.MetricStorage.Uint64Data) == 1 {
					for k, v := range data.MetricStorage.Uint64Data {
						storage.GlobalMetricStorage[userUUID][idx].Uint64Data[k] = v
					}
				}
				return
			}
		}

		storage.GlobalMetricStorage[userUUID] = append(storage.GlobalMetricStorage[userUUID], data.MetricStorage)
	} else {
		storage.GlobalMetricStorage[userUUID] = []metrics.MetricStorage{data.MetricStorage}
	}
	return
}

func (server *Server) UploadMetrics(ctx *fiber.Ctx) (err error) {
	var request clientHTTP.UserMetricStorage

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	unloadUserMetricStorageToGlobalMetricStorage(request)
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
