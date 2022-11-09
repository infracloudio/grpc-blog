package app

import (
	"fmt"
	"github.com/infracloudio/grpc-blog/greet_app/internal/app/server"
	greetpb "github.com/infracloudio/grpc-blog/greet_app/internal/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
)

const (
	SSL_CERT_PATH = "internal/ssl/server.crt"
	SSL_KEY_PATH  = "internal/ssl/server.pem"
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

	certFile := SSL_CERT_PATH
	keyFile := SSL_KEY_PATH

	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalf("Failed to load certificates: %v", sslErr)
		return
	}

	opts = append(opts, grpc.Creds(creds))

	a.grpcServer = grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(a.grpcServer, server.GetNewGreetServer())

	go func() {
		if err := a.grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func (a *App) Shutdown() {
	a.grpcServer.GracefulStop()
}
