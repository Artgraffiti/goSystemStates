package server

import (
	"GSS/internal/client"
	"GSS/internal/metrics"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	pb "GSS/proto"
)

func (server *Server) UploadMetrics(ctx *fiber.Ctx) (err error) {
	var userMetricStorage client.UserMetricStorage

	if err := ctx.BodyParser(&userMetricStorage); err != nil {
		return err
	}

	userUUID := userMetricStorage.UUID
	_, ok := storage[userUUID]
	if ok {
		timestamp := userMetricStorage.MetricStorage.Timestamp
		for idx, metric := range storage[userUUID] {
			if metric.Timestamp == timestamp {
				if len(userMetricStorage.MetricStorage.MetricMapFloat64) == 1 {
					for k, v := range userMetricStorage.MetricStorage.MetricMapFloat64 {
						storage[userUUID][idx].MetricMapFloat64[k] = v
					}
				} else if len(userMetricStorage.MetricStorage.MetricMapUint32) == 1 {
					for k, v := range userMetricStorage.MetricStorage.MetricMapUint32 {
						storage[userUUID][idx].MetricMapUint32[k] = v
					}
				} else if len(userMetricStorage.MetricStorage.MetricMapUint64) == 1 {
					for k, v := range userMetricStorage.MetricStorage.MetricMapUint64 {
						storage[userUUID][idx].MetricMapUint64[k] = v
					}
				}
				return
			}
		}

		storage[userUUID] = append(storage[userUUID], userMetricStorage.MetricStorage)
	} else {
		storage[userUUID] = []metrics.MetricStorage{userMetricStorage.MetricStorage}
	}

	return
}

func (server *Server) MetricsByUUID(ctx *fiber.Ctx) (err error) {
	user_uuid, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return
	}

	return ctx.JSON(storage[user_uuid])
}

func (server *Server) GetUsersMetrics(ctx *fiber.Ctx) (err error) {
	return ctx.JSON(storage)
}

/* ------------------------GRPS SERVER------------------------ */

func (server *GRPCServer) UploadMetrics(ctx context.Context, protoUserMetricStorage *pb.UserMetricStorage) (response *pb.Empty, err error) {
	log.Println("GRPC method UploadMetrics...")
	response = &pb.Empty{}

	userUUID, err := uuid.Parse(protoUserMetricStorage.UUID)
	mStorage := metrics.ConvertProtoToMetricStorage(protoUserMetricStorage.MetricStorage)
	if err != nil {
		return
	}
	_, ok := storage[userUUID]
	if ok {
		timestamp := protoUserMetricStorage.MetricStorage.Timestamp
		for idx, metric := range storage[userUUID] {
			if metric.Timestamp == timestamp {
				if len(protoUserMetricStorage.MetricStorage.MetricMapFloat64) == 1 {
					for k, v := range protoUserMetricStorage.MetricStorage.MetricMapFloat64 {
						storage[userUUID][idx].MetricMapFloat64[k] = v
					}
				} else if len(protoUserMetricStorage.MetricStorage.MetricMapUint32) == 1 {
					for k, v := range protoUserMetricStorage.MetricStorage.MetricMapUint32 {
						storage[userUUID][idx].MetricMapUint32[k] = v
					}
				} else if len(protoUserMetricStorage.MetricStorage.MetricMapUint64) == 1 {
					for k, v := range protoUserMetricStorage.MetricStorage.MetricMapUint64 {
						storage[userUUID][idx].MetricMapUint64[k] = v
					}
				}
				return
			}
		}
		storage[userUUID] = append(storage[userUUID], *mStorage)
	} else {
		storage[userUUID] = []metrics.MetricStorage{*mStorage}
	}
	return
}

func (server *GRPCServer) GetUsersMetrics(ctx context.Context, _ *pb.Empty) (protoUserMetrics *pb.UsersMetrics, err error) {
	log.Println("GRPC method GetUsersMetrics...")
	protoUserMetrics = &pb.UsersMetrics{
		Data: map[string]*pb.MetricStorageArr{},
	}
	for userUUID, metricStorageArr := range storage {
		userUUIDStr := userUUID.String()
		protoUserMetrics.Data[userUUIDStr] = &pb.MetricStorageArr{
			DataArr: []*pb.MetricStorage{},
		}
		for _, metricStorage := range metricStorageArr {
			protoMetricStorage := metrics.ConvertMetricStorageToProto(&metricStorage)
			protoUserMetrics.Data[userUUIDStr].DataArr = append(protoUserMetrics.Data[userUUIDStr].DataArr, protoMetricStorage)
		}
	}
	return
}
