package server

import (
	"context"

	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"github.com/yujisoyama/go_microservices/services/dbmanager/internal/validation"
	"github.com/yujisoyama/go_microservices/services/dbmanager/server/handler"
	"google.golang.org/grpc/codes"
)

func (dbm *DbManager) UpsertUser(ctx context.Context, req *dbmanager.UpsertUserRequest) (*dbmanager.UpsertUserResponse, error) {
	err := validation.ValidateUpsertUserRequest(req)
	if err != nil {
		return nil, utils.GrpcException(codes.InvalidArgument, err.Error())
	}
	
	resp, err := handler.UpsertUser(ctx, dbm.dbClient, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
