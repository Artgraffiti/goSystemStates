package grpc

import (
	"GSS/internal/server/config"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "GSS/proto"
)

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
