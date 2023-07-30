package http

func (server Server) SetupRoutes() {
	server.App.Get("/", server.GetUsersMetrics)
	server.App.Post("/", server.UploadMetrics)

	server.App.Get("/:uuid", server.MetricsByUUID)
}
