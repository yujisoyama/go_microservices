package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty"`
}
