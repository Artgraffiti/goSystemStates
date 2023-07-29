package server

import (
	"GSS/internal/metrics"
	"GSS/internal/server/config"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "GSS/proto"
)

type UsersMetrics map[uuid.UUID][]metrics.MetricStorage

var storage = UsersMetrics{}

type Server struct {
	App    *fiber.App
	Config config.Config
}

func NewServer(config config.Config) (server *Server) {
	server = &Server{
		App:    fiber.New(),
		Config: config,
	}
	server.App.Use(logger.New(logger.Config{
		Format: config.Logger_fmt,
	}))

	server.SetupRoutes()

	return
}

func (server *Server) Run() {
	err := server.App.Listen(server.Config.ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
}

/* ------------------------GRPS SERVER------------------------ */

type GRPCServer struct {
	inner_server *grpc.Server
	Config       config.Config
	pb.UnimplementedGoSystemStatesServer
}

func NewGRPCServer(config config.Config) (serverGRPC *GRPCServer) {
	serverGRPC = &GRPCServer{
		inner_server: grpc.NewServer(
			grpc.Creds(insecure.NewCredentials()),
		),
		Config: config,
	}
	return
}

func (server *GRPCServer) Run() (err error) {
	lis, err := net.Listen("tcp", server.Config.GRPCServerAddr)
	if err != nil {
		return
	}

	pb.RegisterGoSystemStatesServer(server.inner_server, server)

	go func() {
		err = server.inner_server.Serve(lis)
		if err != nil {
			log.Println(err)
		}
	}()
	return
}
