package app

import (
	"fmt"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/infracloudio/grpc-blog/greet_app/internal/app/server"
	"github.com/infracloudio/grpc-blog/greet_app/internal/pkg/middleware"
	greetpb "github.com/infracloudio/grpc-blog/greet_app/internal/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type App struct {
	grpcServer *grpc.Server
}

func GetNewGreetApp() *App {
	return &App{}
}

func (a *App) Start() {
	fmt.Println("Starting greet server on port 50059....")

	servAddr := ":50059"
	if len(os.Getenv("GRPC_SRV_ADDR")) > 0 {
		servAddr = os.Getenv("GRPC_SRV_ADDR")
	}
	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	opts = append(opts, middleware.GetGrpcMiddlewareOpts()...)
	a.grpcServer = grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(a.grpcServer, server.GetNewGreetServer())

	grpc_prometheus.Register(a.grpcServer)
	middleware.RunPrometheusServer()

	go func() {
		if err := a.grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func (a *App) Shutdown() {
	a.grpcServer.GracefulStop()
}
