package storage

import (
	"GSS/internal/metrics"

	"github.com/google/uuid"
)

var GlobalMetricStorage = map[uuid.UUID][]metrics.MetricStorage{}

// TODO: реализовать структуру 'storage' с методом загрузки метрик
