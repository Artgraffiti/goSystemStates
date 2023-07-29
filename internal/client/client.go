package client

import (
	"GSS/internal/client/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "GSS/proto"
)

type User struct {
	UUID   uuid.UUID
	Client *fiber.Client
	Config config.Config
}

func NewUser(config config.Config) (user *User, err error) {
	userUUID, err := uuid.Parse(config.UUID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Init user(uuid): %s", userUUID)
	user = &User{
		UUID:   userUUID,
		Client: fiber.AcquireClient(),
		Config: config,
	}
	return
}

/* ------------------------GRPS CLIENT------------------------ */

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
