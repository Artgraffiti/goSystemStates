package mqtt

import (
	"log"
	"strings"

	"GSS/internal/metrics"
	"GSS/internal/server/storage"
	pb "GSS/proto"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func getLastTopicPart(topic string) string {
	parts := strings.Split(topic, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

func unloadProtoMetricDataToGlobalMetricStorage(data *pb.UserMetricStorage) (err error) {
	userUUID, err := uuid.Parse(data.UUID)
	if err != nil {
		return
	}

	mStorage := metrics.ConvertProtoToMetricStorage(data.Metric)

	_, ok := storage.GlobalMetricStorage[userUUID]
	if ok {
		timestamp := data.Metric.Timestamp
		for idx, metric := range storage.GlobalMetricStorage[userUUID] {
			if metric.Timestamp == timestamp {
				if len(data.Metric.Float64Data) == 1 {
					for k, v := range data.Metric.Float64Data {
						storage.GlobalMetricStorage[userUUID][idx].Float64Data[k] = v
					}
				} else if len(data.Metric.Uint32Data) == 1 {
					for k, v := range data.Metric.Uint32Data {
						storage.GlobalMetricStorage[userUUID][idx].Uint32Data[k] = v
					}
				} else if len(data.Metric.Uint64Data) == 1 {
					for k, v := range data.Metric.Uint64Data {
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

func (server *MQTTServer) UploadMetrics(cli mqtt.Client, msg mqtt.Message) {
	log.Println("MQTT method UploadMetrics...")
	protoUserMetricStorage := pb.UserMetricStorage{}

	if err := proto.Unmarshal(msg.Payload(), &protoUserMetricStorage); err != nil {
		log.Fatal(err)
	}

	unloadProtoMetricDataToGlobalMetricStorage(&protoUserMetricStorage)
}
