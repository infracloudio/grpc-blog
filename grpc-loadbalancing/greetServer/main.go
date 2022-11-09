package main

import (
	"context"
	"fmt"
	greetpb "github.com/infracloudio/grpc-blog/grpc-loadbalancing/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(_ context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	var podIp string
	if len(os.Getenv("POD_IP")) > 0 {
		podIp = os.Getenv("POD_IP")
	}

	result := fmt.Sprintf("Hello, %s %s from %s", firstName, lastName, podIp)
	return &greetpb.GreetingResponse{
		Result: result,
	}, nil
}

func main() {
	fmt.Println("Starting greet server")

	port := "50051"
	if len(os.Getenv("GRPC_PORT")) > 0 {
		port = os.Getenv("GRPC_PORT")
	}

	servAddr := fmt.Sprintf("0.0.0.0:%s", port)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v",err)
	}
}