package http

func (server HTTPServer) SetupRoutes() {
	server.App.Get("/", server.GetUsersMetrics)
	server.App.Post("/", server.UploadMetrics)

	server.App.Get("/:uuid", server.MetricsByUUID)
}
