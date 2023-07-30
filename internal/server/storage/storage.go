package storage

import (
	"GSS/internal/metrics"

	"github.com/google/uuid"
)

var GlobalMetricStorage = map[uuid.UUID][]metrics.MetricStorage{}
