package http

import (
	"GSS/internal/metrics"
	"log"
)

func (user *User) sendUserMetric(url string, request metrics.UserMetricStorage) (err error) {
	agent := user.Client.Post(url)
	err = agent.JSON(request).Parse()
	if err != nil {
		return
	}

	statusCode, _, errs := agent.String()
	log.Println("Send JSON:", statusCode)
	if errs != nil {
		return errs[0]
	}
	return
}

func (user *User) streamingMetricsByFloat64(timestamp int64, mStorage metrics.MetricStorage) (err error) {
	for k, v := range mStorage.Float64Data {
		request := metrics.UserMetricStorage{
			UUID: user.UUID,
			MetricStorage: metrics.MetricStorage{
				Timestamp: timestamp,
				Float64Data: map[string]float64{
					k: v,
				},
				Uint32Data: map[string]uint32{},
				Uint64Data: map[string]uint64{},
			},
		}

		err := user.sendUserMetric("http://"+user.Config.ServerAddr, request)
		if err != nil {
			return err
		}
		log.Printf("Sending metric '%s'...", k)
	}
	return
}

func (user *User) streamingMetricsByUint32(timestamp int64, mStorage metrics.MetricStorage) (err error) {
	for metricKey, metricVal := range mStorage.Uint32Data {
		request := metrics.UserMetricStorage{
			UUID: user.UUID,
			MetricStorage: metrics.MetricStorage{
				Timestamp:   timestamp,
				Float64Data: map[string]float64{},
				Uint32Data: map[string]uint32{
					metricKey: metricVal,
				},
				Uint64Data: map[string]uint64{},
			},
		}

		err := user.sendUserMetric("http://"+user.Config.ServerAddr, request)
		if err != nil {
			return err
		}
		log.Printf("Sending metric '%s'...", metricKey)
	}
	return
}

func (user *User) streamingMetricsByUint64(timestamp int64, mStorage metrics.MetricStorage) (err error) {
	for metricKey, metricVal := range mStorage.Uint64Data {
		request := metrics.UserMetricStorage{
			UUID: user.UUID,
			MetricStorage: metrics.MetricStorage{
				Timestamp:   timestamp,
				Float64Data: map[string]float64{},
				Uint32Data:  map[string]uint32{},
				Uint64Data: map[string]uint64{
					metricKey: metricVal,
				},
			},
		}

		err := user.sendUserMetric("http://"+user.Config.ServerAddr, request)
		if err != nil {
			return err
		}
		log.Printf("Sending metric '%s'...", metricKey)
	}
	return
}

func (user *User) SendMetricStorage() (err error) {
	mStorage, err := metrics.Get()
	if err != nil {
		return
	}

	request := metrics.UserMetricStorage{
		UUID:          user.UUID,
		MetricStorage: mStorage,
	}
	err = user.sendUserMetric("http://"+user.Config.ServerAddr, request)
	if err != nil {
		return
	}

	log.Printf("Sending metrics...")
	return
}

func (user *User) StreamingMetrics() (err error) {
	mStorage, err := metrics.Get()
	if err != nil {
		return
	}
	timestamp := mStorage.Timestamp

	user.streamingMetricsByFloat64(timestamp, mStorage)
	user.streamingMetricsByUint32(timestamp, mStorage)
	user.streamingMetricsByUint64(timestamp, mStorage)
	return
}
