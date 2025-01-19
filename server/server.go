package main

import (
	"context"
	"log"
	"net"

	pb "github.com/joaopaulo-bertoncini/grpc-test/proto"

	"google.golang.org/grpc"
)

// Implementação do serviço Greeter
type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + req.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &greeterServer{})

	log.Println("Server is running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
