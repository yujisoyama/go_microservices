package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/yujisoyama/go_microservices/services/dbmanager/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DB         = "public"
	COLLECTION = "users"
)

func UpsertUser(ctx context.Context, dbClient *mongo.Client, newUser *entity.UserEntity) (*entity.UserEntity, error) {
	session, err := dbClient.StartSession()
	if err != nil {
		return nil, fmt.Errorf("Error in starting session: %v", err)
	}
	defer session.EndSession(ctx)

	resp := &entity.UserEntity{}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		collection := dbClient.Database(DB).Collection(COLLECTION)

		var usr entity.UserEntity
		err = collection.FindOne(ctx, bson.M{"oauth_id": newUser.OauthId}).Decode(&usr)
		if err == mongo.ErrNoDocuments {
			newUser.CreatedAt = time.Now()
			newUser.UpdatedAt = time.Now()
			res, err := collection.InsertOne(ctx, newUser)
			if err != nil {
				return nil, fmt.Errorf("Error in insert user: %v", err)
			}

			oId, ok := res.InsertedID.(primitive.ObjectID)
			if !ok {
				return nil, fmt.Errorf("Error in converting ID of insert user: %v", err)
			}

			newUser.ID = oId
			resp = newUser
		} else {
			newUser.CreatedAt = usr.CreatedAt
			newUser.UpdatedAt = time.Now()
			update := bson.M{
				"$set": newUser,
			}

			_, err = collection.UpdateOne(ctx, bson.M{"_id": usr.ID}, update)
			if err != nil {
				return nil, fmt.Errorf("Error in updating user: %v", err)
			}

			newUser.ID = usr.ID
			resp = newUser
		}

		return nil, nil
	}
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return nil, fmt.Errorf("Error in transaction: %v", err)
	}

	return resp, nil
}

func GetUserById(ctx context.Context, dbClient *mongo.Client, id string) (*entity.UserEntity, error) {
	collection := dbClient.Database(DB).Collection(COLLECTION)
	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("User id is not a valid ObjectID: %v", err)
	}

	var usr entity.UserEntity
	err = collection.FindOne(ctx, bson.M{"_id": oId}).Decode(&usr)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("User not found: %v", err)
	}

	return &usr, nil
}
