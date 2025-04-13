package grpc

import (
	"context"
	"time"

	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func apiKeyUnaryInterceptor(apiKey string) grpc.UnaryClientInterceptor {
	return func(
			ctx context.Context,
			method string,
			req, reply interface{},
			cc *grpc.ClientConn,
			invoker grpc.UnaryInvoker,
			opts ...grpc.CallOption,
	) error {
			// Adiciona o metadata com a API Key
			md := metadata.Pairs("api-key", apiKey)
			ctx = metadata.NewOutgoingContext(ctx, md)
			return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func InitGrpcClient(logger *logger.Logger, dbmHost string, dbmApikey string) (dbmanager.DbManagerClient, error) {
	conn, err := grpc.NewClient(
		dbmHost, 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(apiKeyUnaryInterceptor(dbmApikey)),
	)
	if err != nil {
		logger.Error("Failed to create grpc client", err)
		return nil, err
	}
	client := dbmanager.NewDbManagerClient(conn)
	md := metadata.New(map[string]string{
		"api-key": dbmApikey,
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err = client.Ping(ctx, &dbmanager.PingRequest{})
	if err != nil {
		return nil, err
	}

	logger.Info("Successful connection with DBManager gRPC server!")

	return client, nil
}
