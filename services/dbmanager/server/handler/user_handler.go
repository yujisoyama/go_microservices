package handler

import (
	"context"
	"time"

	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
	"github.com/yujisoyama/go_microservices/pkg/pb/user"
	"github.com/yujisoyama/go_microservices/services/dbmanager/server/dto"
	"github.com/yujisoyama/go_microservices/services/dbmanager/server/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpsertUser(ctx context.Context, dbClient *mongo.Client, req *dbmanager.UpsertUserRequest) (*dbmanager.UpsertUserResponse, error) {
	newUser := dbmanagerdto.InputDtoUserToEntity(req)
	updatedUser, err := repository.UpsertUser(ctx, dbClient, newUser)
	if err != nil {
		return nil, err
	}

	resp := &dbmanager.UpsertUserResponse{
		User: &user.User{
			Id:            updatedUser.ID.Hex(),
			OauthId:       updatedUser.OauthId,
			OauthType:     updatedUser.OauthType,
			Email:         updatedUser.Email,
			VerifiedEmail: updatedUser.VerifiedEmail,
			FirstName:     updatedUser.FirstName,
			LastName:      updatedUser.LastName,
			Picture:       updatedUser.Picture,
			CreatedAt:     updatedUser.CreatedAt.UTC().Format(time.RFC3339),
			UpdatedAt:     updatedUser.UpdatedAt.UTC().Format(time.RFC3339),
		},
	}
	return resp, nil
}
