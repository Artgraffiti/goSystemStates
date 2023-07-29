package server

import (
	"GSS/internal/client"
	"GSS/internal/metrics"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	userUUID, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return
	}

	return ctx.JSON(storage[userUUID])
}

func (server *Server) GetUsersMetrics(ctx *fiber.Ctx) (err error) {
	return ctx.JSON(storage)
}

/* ------------------------GRPS SERVER------------------------ */

func (server *GRPCServer) UploadMetrics(ctx context.Context, protoUserMetricStorage *pb.UserMetricStorage) (response *pb.Empty, err error) {
	log.Println("GRPC method UploadMetrics...")
	response = &pb.Empty{}

	userUUID, err := uuid.Parse(protoUserMetricStorage.UUID)
	if err != nil {
		return
	}

	if protoUserMetricStorage.MetricStorage == nil {
		return nil, status.Error(codes.InvalidArgument, "metric storage is not set")
	}
	mStorage := metrics.ConvertProtoToMetricStorage(protoUserMetricStorage.MetricStorage)

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

func (server *GRPCServer) MetricsByUUID(ctx context.Context, requset *pb.MetricsByUUIDRequest) (response *pb.MetricsByUUIDResponse, err error) {
	log.Println("GRPC method MetricsByUUID...")
	userUUID, err := uuid.Parse(requset.UUID)
	if err != nil {
		return
	}
	response = &pb.MetricsByUUIDResponse{}
	for _, metric := range storage[userUUID] {
		response.MetricsArray = append(response.MetricsArray, metrics.ConvertMetricStorageToProto(&metric))
	}
	return response, err
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
