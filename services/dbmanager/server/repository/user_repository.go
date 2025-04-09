package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
	userpb "github.com/yujisoyama/go_microservices/pkg/pb/user"
	"github.com/yujisoyama/go_microservices/services/dbmanager/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DB         = "public"
	COLLECTION = "users"
)

func UpsertUser(ctx context.Context, dbClient *mongo.Client, user *dbmanager.UpsertUserRequest) (*dbmanager.UpsertUserResponse, error) {
	session, err := dbClient.StartSession()
	if err != nil {
		return nil, fmt.Errorf("Error in starting session: %v", err)
	}
	defer session.EndSession(ctx)

	resp := &dbmanager.UpsertUserResponse{}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		collection := dbClient.Database(DB).Collection(COLLECTION)

		var usr entity.UserEntity
		err = collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&usr)
		if err == mongo.ErrNoDocuments {
			newUser := &entity.UserEntity{
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			res, err := collection.InsertOne(ctx, newUser)
			if err != nil {
				return nil, fmt.Errorf("Error in insert user: %v", err)
			}

			oId, ok := res.InsertedID.(primitive.ObjectID)
			if !ok {
				return nil, fmt.Errorf("Error in converting ID of insert user: %v", err)
			}

			resp.User = &userpb.User{
				Id:        oId.Hex(),
				Name:      newUser.Name,
				Email:     newUser.Email,
				CreatedAt: newUser.CreatedAt.UTC().Format(time.RFC3339),
				UpdatedAt: newUser.UpdatedAt.UTC().Format(time.RFC3339),
			}
			resp.User.Id = oId.Hex()
		} else {
			upUser := &entity.UserEntity{
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: usr.CreatedAt,
				UpdatedAt: time.Now(),
			}

			update := bson.M{
				"$set": upUser,
			}

			_, err = collection.UpdateOne(ctx, bson.M{"_id": usr.ID}, update)
			if err != nil {
				return nil, fmt.Errorf("Error in updating user: %v", err)
			}

			resp.User = &userpb.User{
				Id:        usr.ID.Hex(),
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: usr.CreatedAt.UTC().Format(time.RFC3339),
				UpdatedAt: upUser.UpdatedAt.UTC().Format(time.RFC3339),
			}
		}

		return nil, nil
	}
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return nil, fmt.Errorf("Error in transaction: %v", err)
	}

	return resp, nil
}
