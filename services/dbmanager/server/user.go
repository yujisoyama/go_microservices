package server

import (
	"context"

	"github.com/yujisoyama/go_microservices/pkg/protos/dbmanager"
	"github.com/yujisoyama/go_microservices/pkg/protos/user"
	"github.com/yujisoyama/go_microservices/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
)

func (dbm *DbManager) UpsertUser(ctx context.Context, req *dbmanager.UpsertUserRequest) (*dbmanager.UpsertUserResponse, error) {
	collection := dbm.dbClient.Database("public").Collection("users")

	resp := &dbmanager.UpsertUserResponse{}

	var usr bson.M
	err := collection.FindOne(ctx, bson.M{"email": req.User.Email}).Decode(&usr)
	if err == mongo.ErrNoDocuments {
		dbm.log.Info("User not found, inserting new user")
		newUser := &user.User{
			Name:  req.User.Name,
			Email: req.User.Email,
		}

		res, err := collection.InsertOne(ctx, newUser)
		if err != nil {
			return nil, utils.GrpcException(codes.Internal, "Error in insert user", err)
		}

		oId, ok := res.InsertedID.(primitive.ObjectID)
		if !ok {
			return nil, utils.GrpcException(codes.Internal, "Error in converting ID of insert user", err)
		}

		resp.User = newUser
		resp.User.Id = oId.Hex()
	} else {
		dbm.log.Infof("User with email %s founded. Upserting data", req.User.Email)

		oId, ok := usr["_id"].(primitive.ObjectID)
		if !ok {
			return nil, utils.GrpcException(codes.Internal, "Error in getting ID of user", err)
		}

		update := bson.M{
			"$set": bson.M{
				"name":  req.User.Name,
				"email": req.User.Email,
			},
		}

		_, err = collection.UpdateOne(ctx, bson.M{"_id": oId}, update)
		if err != nil {
			return nil, utils.GrpcException(codes.Internal, "Error in updating user", err)
		}

		resp.User = &user.User{
			Id:    oId.Hex(),
			Name:  req.User.Name,
			Email: req.User.Email,
		}
	}

	return resp, nil
}
