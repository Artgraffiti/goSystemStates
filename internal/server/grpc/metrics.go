package grpc

import (
	"GSS/internal/metrics"
	"GSS/internal/server/storage"
	"context"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "GSS/proto"
)

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

	_, ok := storage.GlobalMetricStorage[userUUID]
	if ok {
		timestamp := protoUserMetricStorage.Metric.Timestamp
		for idx, metric := range storage.GlobalMetricStorage[userUUID] {
			if metric.Timestamp == timestamp {
				if len(protoUserMetricStorage.Metric.Float64Data) == 1 {
					for k, v := range protoUserMetricStorage.Metric.Float64Data {
						storage.GlobalMetricStorage[userUUID][idx].Float64Data[k] = v
					}
				} else if len(protoUserMetricStorage.Metric.Uint32Data) == 1 {
					for k, v := range protoUserMetricStorage.Metric.Uint32Data {
						storage.GlobalMetricStorage[userUUID][idx].Uint32Data[k] = v
					}
				} else if len(protoUserMetricStorage.Metric.Uint64Data) == 1 {
					for k, v := range protoUserMetricStorage.Metric.Uint64Data {
						storage.GlobalMetricStorage[userUUID][idx].Uint64Data[k] = v
					}
				}
				return
			}
		}
		storage.GlobalMetricStorage[userUUID] = append(storage.GlobalMetricStorage[userUUID], *mStorage)
	} else {
		storage.GlobalMetricStorage[userUUID] = []metrics.MetricStorage{*mStorage}
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
	for _, metric := range storage.GlobalMetricStorage[userUUID] {
		response.DataArr = append(response.DataArr, metrics.ConvertMetricStorageToProto(&metric))
	}
	return response, err
}

func (server *GRPCServer) GetUsersMetrics(ctx context.Context, _ *pb.Empty) (protoUserMetrics *pb.UsersMetrics, err error) {
	log.Println("GRPC method GetUsersMetrics...")
	protoUserMetrics = &pb.UsersMetrics{
		Data: map[string]*pb.MetricStorageArr{},
	}
	for userUUID, metricStorageArr := range storage.GlobalMetricStorage {
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
