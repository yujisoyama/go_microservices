package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	OauthId       string             `bson:"oauth_id"`
	OauthType     string             `bson:"oauth_type"`
	Email         string             `bson:"email"`
	VerifiedEmail bool               `bson:"verified_email"`
	FirstName     string             `bson:"first_name"`
	LastName      string             `bson:"last_name"`
	Picture       string             `bson:"picture"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
	DeletedAt     time.Time          `bson:"deleted_at,omitempty"`
}
