package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/my-name/grpc-service-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "localhost:5050"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewGeometryServiceClient(conn)

	area, err := grpcClient.Area(ctx, &pb.RectRequest{Height: 10.1, Width: 20.5})
	if err != nil {
		log.Fatalf("Area failed: %v", err)
	}

	perim, err := grpcClient.Perimeter(ctx, &pb.RectRequest{Height: 10.1, Width: 20.5})
	if err != nil {
		log.Fatalf("Perimeter failed: %v", err)
	}

	fmt.Printf("Area: %.2f\nPerimeter: %.2f\n", area.Result, perim.Result)
}
