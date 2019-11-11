package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Person struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Enabled   bool               `bson:"enabled"`
	Type      int32              `bson:"type"`
	Name      string             `bson:"name"`
	CreatedAt *time.Time         `bson:"created_at,omitempty"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty"`
}
