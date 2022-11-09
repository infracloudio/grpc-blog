package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/infracloudio/grpc-blog/greet_app/internal/app/server"
	greetpb "github.com/infracloudio/grpc-blog/greet_app/internal/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"strings"
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

	cert, err := tls.LoadX509KeyPair("internal/app/server/cert/public.crt", "internal/app/server/cert/private.key")
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	opts := []grpc.ServerOption{grpc.UnaryInterceptor(validateToken), grpc.Creds(credentials.NewServerTLSFromCert(&cert))}
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

func validateToken(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	if !valid(md["authorization"]) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return handler(ctx, req)
}

func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}

	token := strings.TrimPrefix(authorization[0], "Bearer ")

	return token == "client-x-id"
}
