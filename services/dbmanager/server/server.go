package server

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/pkg/protos/dbmanager"
	"github.com/yujisoyama/go_microservices/services/dbmanager/internal/interceptor"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type DbManager struct {
	dbmanager.UnimplementedDbManagerServer
	log      *logrus.Logger
	configs  *DbManagerConfigs
	dbClient *mongo.Client
}

func NewDbManager() *DbManager {
	return &DbManager{
		log:     logger.NewLogger(),
		configs: &DbManagerConfigs{},
	}
}

func (dbm *DbManager) Run(ctx context.Context) error {
	dbm.SetConfigs()
	dbm.log.Info("Start grpc dbmanager in port: ", dbm.configs.port)

	clientOptions := options.Client().ApplyURI(dbm.DbConnectString())

	var err error
	dbm.dbClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer dbm.dbClient.Disconnect(ctx)

	err = dbm.dbClient.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	dbm.log.Info("Connected to MongoDB!")

	uInterceptors := grpc.ChainUnaryInterceptor(
		interceptor.LoggingInterceptor(dbm.log),
		interceptor.AuthInterceptor(dbm.log, dbm.configs.apikey),
	)

	grpcServer := grpc.NewServer(uInterceptors)
	reflection.Register(grpcServer)
	dbmanager.RegisterDbManagerServer(grpcServer, dbm)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", dbm.configs.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	dbm.log.Infof("Ready to serve requests!")
	if err = grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
