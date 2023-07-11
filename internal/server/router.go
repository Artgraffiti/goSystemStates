package server

func (server Server) SetupRoutes() {
	server.App.Get("/", server.Hi)
	server.App.Post("/", server.addMetric)

	server.App.Get("/metrics", server.getAllMetrics)
}
