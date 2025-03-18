package server

import (
	"context"
	"time"

	"github.com/yujisoyama/go_microservices/pkg/protos/dbmanager"
	"github.com/yujisoyama/go_microservices/pkg/protos/user"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"github.com/yujisoyama/go_microservices/services/dbmanager/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
)

func (dbm *DbManager) UpsertUser(ctx context.Context, req *dbmanager.UpsertUserRequest) (*dbmanager.UpsertUserResponse, error) {
	collection := dbm.dbClient.Database("public").Collection("users")

	resp := &dbmanager.UpsertUserResponse{}

	var usr entity.UserEntity
	err := collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&usr)
	if err == mongo.ErrNoDocuments {
		dbm.log.Info("User not found, inserting new user")
		newUser := &entity.UserEntity{
			Name:      req.Name,
			Email:     req.Email,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		res, err := collection.InsertOne(ctx, newUser)
		if err != nil {
			return nil, utils.GrpcException(codes.Internal, "Error in insert user", err)
		}

		oId, ok := res.InsertedID.(primitive.ObjectID)
		if !ok {
			return nil, utils.GrpcException(codes.Internal, "Error in converting ID of insert user", err)
		}

		resp.User = &user.User{
			Id:        oId.Hex(),
			Name:      newUser.Name,
			Email:     newUser.Email,
			CreatedAt: newUser.CreatedAt.UTC().Format(time.RFC3339),
			UpdatedAt: newUser.UpdatedAt.UTC().Format(time.RFC3339),
		}
		resp.User.Id = oId.Hex()
	} else {
		dbm.log.Infof("User with email %s founded. Upserting data", req.Email)

		upUser := &entity.UserEntity{
			Name:      req.Name,
			Email:     req.Email,
			CreatedAt: usr.CreatedAt,
			UpdatedAt: time.Now(),
		}

		update := bson.M{
			"$set": upUser,
	}

		_, err = collection.UpdateOne(ctx, bson.M{"_id": usr.ID}, update)
		if err != nil {
			return nil, utils.GrpcException(codes.Internal, "Error in updating user", err)
		}

		resp.User = &user.User{
			Id:    usr.ID.Hex(),
			Name:  req.Name,
			Email: req.Email,
			CreatedAt: usr.CreatedAt.UTC().Format(time.RFC3339),
			UpdatedAt: upUser.UpdatedAt.UTC().Format(time.RFC3339),
		}
	}

	return resp, nil
}
