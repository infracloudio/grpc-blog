package middleware

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func AddInterceptors(opts []grpc.ServerOption, uInterceptors []grpc.UnaryServerInterceptor, sInterceptors []grpc.StreamServerInterceptor) []grpc.ServerOption {
	opts = append(opts, grpc.ChainUnaryInterceptor(uInterceptors...))
	opts = append(opts, grpc.ChainStreamInterceptor(sInterceptors...))
	return opts
}

func AddPrometheus(uInterceptors *[]grpc.UnaryServerInterceptor, sInterceptors *[]grpc.StreamServerInterceptor) {
	*uInterceptors = append(*uInterceptors, grpc_prometheus.UnaryServerInterceptor)
	*sInterceptors = append(*sInterceptors, grpc_prometheus.StreamServerInterceptor)
}

// Add grpc default middleware like logging and prometheus metrics
func GetGrpcMiddlewareOpts() []grpc.ServerOption {
	// gRPC server startup options
	opts := []grpc.ServerOption{}

	uInterceptors := []grpc.UnaryServerInterceptor{}
	sInterceptors := []grpc.StreamServerInterceptor{}

	// add middleware
	AddLogging(&zap.Logger{}, &uInterceptors, &sInterceptors)
	AddPrometheus(&uInterceptors, &sInterceptors)

	opts = AddInterceptors(opts, uInterceptors, sInterceptors)

	return opts
}
