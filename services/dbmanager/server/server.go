package server

import (
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/yujisoyama/go_microservices/pkg/logger"
	"github.com/yujisoyama/go_microservices/pkg/protos/dbmanager"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type DbManager struct {
	dbmanager.UnimplementedDbManagerServer
	port string
	apikey string
	log *logrus.Logger
}

func NewDbManager() *DbManager {
	return &DbManager{
		log: logger.NewLogger(),
	}
}

func (dbM *DbManager) SetConfigs() {
	dbM.port = utils.GetEnv("PORT")
	dbM.apikey = utils.GetEnv("API_KEY")
}

func (dbM *DbManager) Run(ctx context.Context) error {
	dbM.SetConfigs()
	dbM.log.Info("Start grpc dbmanager in port: ", dbM.port)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	dbmanager.RegisterDbManagerServer(grpcServer, dbM)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", dbM.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	dbM.log.Infof("Ready to serve requests!")
	if err = grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}