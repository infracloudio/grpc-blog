package middleware

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// codeToLevel redirects OK to DEBUG level logging instead of Info
func codeToLevel(code codes.Code) zapcore.Level {
	if code == codes.OK {
		return zap.DebugLevel
	}
	return grpc_zap.DefaultCodeToLevel(code)
}

func extractFields(fullMethod string, req interface{}) map[string]interface{} {
	return make(map[string]interface{})
}

// AddLogging returns grpc.Server config option that turn on logging.
func AddLogging(logger *zap.Logger, uInterceptors *[]grpc.UnaryServerInterceptor, sInterceptors *[]grpc.StreamServerInterceptor) {
	// Shared options for the logger, with a custom gRPC code to log level function.
	o := []grpc_zap.Option{
		grpc_zap.WithLevels(codeToLevel),
	}

	// Make sure that log statements internal to gRPC library are logged using the zaplogger as well.
	grpc_zap.ReplaceGrpcLoggerV2(logger)

	*uInterceptors = append(*uInterceptors, grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(extractFields)))
	*uInterceptors = append(*uInterceptors, grpc_zap.UnaryServerInterceptor(logger, o...))

	*sInterceptors = append(*sInterceptors, grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)))
	*sInterceptors = append(*sInterceptors, grpc_zap.StreamServerInterceptor(logger, o...))

}
