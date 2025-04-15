package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/my-name/grpc-service-example/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GeometryServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Area(ctx context.Context, in *pb.RectRequest) (*pb.AreaResponse, error) {
	log.Println("invoked Area: ", in)
	return &pb.AreaResponse{Result: in.Height * in.Width}, nil
}

func (s *Server) Perimeter(ctx context.Context, in *pb.RectRequest) (*pb.PerimeterResponse, error) {
	log.Println("invoked Perimeter: ", in)
	return &pb.PerimeterResponse{Result: 2 * (in.Height + in.Width)}, nil
}

func main() {
	host := "0.0.0.0"
	port := "5050"

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	grpcServer := grpc.NewServer()
	pb.RegisterGeometryServiceServer(grpcServer, NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
