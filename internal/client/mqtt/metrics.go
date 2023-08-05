package mqtt

import (
	"GSS/internal/metrics"
	"log"

	"google.golang.org/protobuf/proto"
)

func (user *MQTTUser) UploadMetrics() (err error) {
	mStorage, err := metrics.Get()
	if err != nil {
		return
	}

	umStorage := metrics.UserMetricStorage{
		UUID:          user.UUID,
		MetricStorage: mStorage,
	}

	protoMStorage := metrics.ConvertUserMetricStorageToProto(&umStorage)
	protoBytes, err := proto.Marshal(protoMStorage)
	path := "goSystemStates/metrics/" + user.UUID.String()

	log.Println("Send metrics(MQTT)...")
	if token := user.client.Publish(path, 0, false, protoBytes); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return
}
