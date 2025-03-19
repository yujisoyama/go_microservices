package interceptor

import (
	"context"

	"github.com/yujisoyama/go_microservices/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(l *logger.Logger, apiKey string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, h grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			l.Errorf("Empty metadata")
			return nil, status.Error(codes.Unauthenticated, "Empty metadata")
		}

		// Obtém a chave "api-key" do metadata
		reqKey := md.Get("api-key")
		if apiKey != reqKey[0] {
			l.Errorf("Invalid api key")
			return nil, status.Error(codes.Unauthenticated, "Invalid api key")
		}

		result, err := h(ctx, req)

		return result, err
	}
}
