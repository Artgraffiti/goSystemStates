package mqtt

func (server *MQTTServer) SetupRoutes() {
	server.client.Subscribe("goSystemStates/metrics/+", 0, server.UploadMetrics)
}
