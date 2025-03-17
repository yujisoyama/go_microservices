package server

import (
	"context"

	"github.com/yujisoyama/go_microservices/pkg/protos/dbmanager"
)

func (dbM *DbManager) CreateUser(ctx context.Context, in *dbmanager.CreateUserRequest) (*dbmanager.CreateUserResponse, error) {
	dbM.log.Info("CreateUser request")
	return &dbmanager.CreateUserResponse{}, nil
}