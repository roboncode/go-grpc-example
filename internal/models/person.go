package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Person struct {
	Id        *primitive.ObjectID `bson:"_id,omitempty"`
	Name      string              `bson:"name"`
	Email     string              `bson:"email"`
	Status    int32               `bson:"status"`
	CreatedAt *time.Time          `bson:"created_at,omitempty"`
	UpdatedAt *time.Time          `bson:"updated_at,omitempty"`
}
