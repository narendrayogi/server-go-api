package v1

import (
	"context"
	"log"
	"time"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gofiber/fiber/v3"
	"github.com/narendrayogi/server-go-api/handler/v1/product"
)

func ListProductHandler(
	c fiber.Ctx,
) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := product.NewProductServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	isActive := true
	r, err := client.ListProduct(ctx, &product.ListProductReq{
		IsActive: &isActive,
	})
	if err != nil {
		log.Fatalf("could not get list: %v", err)
	}
	u, err := json.Marshal(r);

        if err != nil {
            panic(err)
        }

	// return c.SendString(r)
	return c.Send([]byte(string(u)))
}