package server

import (
	"context"

	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
)

func (dbM *DbManager) Ping(ctx context.Context, in *dbmanager.PingRequest) (*dbmanager.PingResponse, error) {
	dbM.log.Info("Ping")
	return &dbmanager.PingResponse{}, nil
}
