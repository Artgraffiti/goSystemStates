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
	_, ok := GlobalMetricStorage[userUUID]
	if ok {
		timestamp := userMetricStorage.MetricStorage.Timestamp
		for idx, metric := range GlobalMetricStorage[userUUID] {
			if metric.Timestamp == timestamp {
				if len(userMetricStorage.MetricStorage.Float64Data) == 1 {
					for k, v := range userMetricStorage.MetricStorage.Float64Data {
						GlobalMetricStorage[userUUID][idx].Float64Data[k] = v
					}
				} else if len(userMetricStorage.MetricStorage.Uint32Data) == 1 {
					for k, v := range userMetricStorage.MetricStorage.Uint32Data {
						GlobalMetricStorage[userUUID][idx].Uint32Data[k] = v
					}
				} else if len(userMetricStorage.MetricStorage.Uint64Data) == 1 {
					for k, v := range userMetricStorage.MetricStorage.Uint64Data {
						GlobalMetricStorage[userUUID][idx].Uint64Data[k] = v
					}
				}
				return
			}
		}

		GlobalMetricStorage[userUUID] = append(GlobalMetricStorage[userUUID], userMetricStorage.MetricStorage)
	} else {
		GlobalMetricStorage[userUUID] = []metrics.MetricStorage{userMetricStorage.MetricStorage}
	}

	return
}

func (server *Server) MetricsByUUID(ctx *fiber.Ctx) (err error) {
	userUUID, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return
	}

	return ctx.JSON(GlobalMetricStorage[userUUID])
}

func (server *Server) GetUsersMetrics(ctx *fiber.Ctx) (err error) {
	return ctx.JSON(GlobalMetricStorage)
}

/* ------------------------GRPS SERVER------------------------ */

func (server *GRPCServer) UploadMetrics(ctx context.Context, protoUserMetricStorage *pb.UserMetricStorage) (response *pb.Empty, err error) {
	log.Println("GRPC method UploadMetrics...")
	response = &pb.Empty{}

	userUUID, err := uuid.Parse(protoUserMetricStorage.UUID)
	if err != nil {
		return
	}

	if protoUserMetricStorage.Metric == nil {
		return nil, status.Error(codes.InvalidArgument, "metric storage is not set")
	}
	mStorage := metrics.ConvertProtoToMetricStorage(protoUserMetricStorage.Metric)

	_, ok := GlobalMetricStorage[userUUID]
	if ok {
		timestamp := protoUserMetricStorage.Metric.Timestamp
		for idx, metric := range GlobalMetricStorage[userUUID] {
			if metric.Timestamp == timestamp {
				if len(protoUserMetricStorage.Metric.Float64Data) == 1 {
					for k, v := range protoUserMetricStorage.Metric.Float64Data {
						GlobalMetricStorage[userUUID][idx].Float64Data[k] = v
					}
				} else if len(protoUserMetricStorage.Metric.Uint32Data) == 1 {
					for k, v := range protoUserMetricStorage.Metric.Uint32Data {
						GlobalMetricStorage[userUUID][idx].Uint32Data[k] = v
					}
				} else if len(protoUserMetricStorage.Metric.Uint64Data) == 1 {
					for k, v := range protoUserMetricStorage.Metric.Uint64Data {
						GlobalMetricStorage[userUUID][idx].Uint64Data[k] = v
					}
				}
				return
			}
		}
		GlobalMetricStorage[userUUID] = append(GlobalMetricStorage[userUUID], *mStorage)
	} else {
		GlobalMetricStorage[userUUID] = []metrics.MetricStorage{*mStorage}
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
	for _, metric := range GlobalMetricStorage[userUUID] {
		response.DataArr = append(response.DataArr, metrics.ConvertMetricStorageToProto(&metric))
	}
	return response, err
}

func (server *GRPCServer) GetUsersMetrics(ctx context.Context, _ *pb.Empty) (protoUserMetrics *pb.UsersMetrics, err error) {
	log.Println("GRPC method GetUsersMetrics...")
	protoUserMetrics = &pb.UsersMetrics{
		Data: map[string]*pb.MetricStorageArr{},
	}
	for userUUID, metricStorageArr := range GlobalMetricStorage {
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
