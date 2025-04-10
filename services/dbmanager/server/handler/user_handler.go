package handler

import (
	"context"

	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
	"github.com/yujisoyama/go_microservices/services/dbmanager/server/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpsertUser(ctx context.Context, dbClient *mongo.Client, user *dbmanager.UpsertUserRequest) (*dbmanager.UpsertUserResponse, error) {
	resp, err := repository.UpsertUser(ctx, dbClient, user)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
