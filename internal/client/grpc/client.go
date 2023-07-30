package grpc

import (
	"GSS/internal/client/config"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "GSS/proto"
)

type GRPCUser struct {
	UUID   uuid.UUID
	Client pb.GoSystemStatesClient
	Config config.Config
}

func NewGRPCUser(config config.Config) (user *GRPCUser, err error) {
	userUUID, err := uuid.Parse(config.UUID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Init GRPCUser(uuid): %s", userUUID)
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	user = &GRPCUser{
		UUID:   userUUID,
		Client: pb.NewGoSystemStatesClient(conn),
		Config: config,
	}
	return
}
