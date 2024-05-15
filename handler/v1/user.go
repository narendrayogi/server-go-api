package v1

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gofiber/fiber/v3"
	"github.com/narendrayogi/server-go-api/handler/v1/user"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)


func ListUserHandler(
	c fiber.Ctx,
) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := user.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	isActive := true
	r, err := client.ListUser(ctx, &user.ListUserReq{
		IsActive: &isActive,
	})
	if err != nil {
		log.Fatalf("could not get list: %v", err)
	}

	// return c.SendString(r)
	return c.Send([]byte(r.String()))
}

