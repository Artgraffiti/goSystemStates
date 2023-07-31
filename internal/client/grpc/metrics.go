package grpc

import (
	"GSS/internal/metrics"
	"context"
	"log"
)

func (user *GRPCUser) UploadMetrics() (err error) {
	mStorage, err := metrics.Get()
	if err != nil {
		return
	}

	umStorage := metrics.UserMetricStorage{
		UUID:          user.UUID,
		MetricStorage: mStorage,
	}

	request := metrics.ConvertUserMetricStorageToProto(&umStorage)
	log.Println("Send metrics(GRPC)...")
	_, err = user.client.UploadMetrics(context.Background(), request)
	return
}
