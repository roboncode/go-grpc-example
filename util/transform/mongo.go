package transform

import "go.mongodb.org/mongo-driver/bson/primitive"

func ToObjectId(id string) *primitive.ObjectID {
	if oid, err := primitive.ObjectIDFromHex(id); err != nil {
		return nil
	} else {
		return &oid
	}
}
