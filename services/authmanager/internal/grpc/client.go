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

func InitGrpcClient(logger *logger.Logger, dbmHost string, dbmApikey string) (dbmanager.DbManagerClient, error) {
	conn, err := grpc.NewClient(dbmHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
