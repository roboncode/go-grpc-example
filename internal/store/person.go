package store

import (
	"context"
	"example/generated"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const PersonStoreName = "Person"

type Person struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Type      int32              `bson:"type"`
	Enabled   bool               `bson:"enabled"`
	CreatedAt *time.Time         `bson:"created_at,omitempty"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty"`
}

type PersonStore struct {
	collection *mongo.Collection
}

func (s *PersonStore) CreatePerson(_ context.Context, req *example.CreatePersonRequest) (*example.Person, error) {
	var now = time.Now()
	var doc = Person{
		Name:      req.GetName(),
		Type:      req.GetType(),
		Enabled:   req.GetEnabled(),
		CreatedAt: &now,
	}
	result, err := s.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	createdAt, err := ptypes.TimestampProto(now)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &example.Person{
		Id:        result.InsertedID.(primitive.ObjectID).Hex(),
		Name:      req.GetName(),
		Enabled:   req.GetEnabled(),
		Type:      req.GetType(),
		CreatedAt: createdAt,
	}, nil
}

func (s *PersonStore) GetPerson(_ context.Context, req *example.GetPersonRequest) (*example.Person, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := s.collection.FindOne(context.Background(), bson.M{"_id": oid})
	// Create an empty BlogItem to write our decode result to
	data := Person{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Person with Object Id %s: %v", req.GetId(), err))
	}
	var createdAt *timestamp.Timestamp
	var updatedAt *timestamp.Timestamp
	createdAt, err = ptypes.TimestampProto(*data.CreatedAt)
	if err != nil {
		return nil, err
	}
	if data.UpdatedAt != nil {
		updatedAt, err = ptypes.TimestampProto(*data.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	response := &example.Person{
		Id:        oid.Hex(),
		Name:      data.Name,
		Type:      data.Type,
		Enabled:   data.Enabled,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return response, nil
}

func (s *PersonStore) GetPersons(_ context.Context, req *example.GetPersonsRequest) (*example.Persons, error) {
	// collection.Find returns a cursor for our (empty) query
	cursor, err := s.collection.Find(context.Background(), bson.M{
		"type":    req.GetType(),
		"enabled": req.GetEnabled(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())

	var items = make([]*example.Person, 0)
	var createdAt *timestamp.Timestamp
	var updatedAt *timestamp.Timestamp
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		var data = &Person{}
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		createdAt, err = ptypes.TimestampProto(*data.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode createdAt: %v", err))
		}
		if data.UpdatedAt != nil {
			updatedAt, err = ptypes.TimestampProto(*data.UpdatedAt)
			if err != nil {
				return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode updatedAt: %v", err))
			}
		}

		items = append(items, &example.Person{
			Id:        data.Id.Hex(),
			Name:      data.Name,
			Type:      data.Type,
			Enabled:   data.Enabled,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return &example.Persons{Items: items}, nil
}

func (s *PersonStore) UpdatePerson(_ context.Context, req *example.UpdatePersonRequest) (*empty.Empty, error) {
	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied blog id to a MongoDB ObjectId: %v", err),
		)
	}

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{
		"name":       req.GetName(),
		"type":       req.GetType(),
		"enabled":    req.GetEnabled(),
		"updated_at": time.Now(),
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := s.collection.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'data'
	data := Person{}
	err = result.Decode(&data)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find Person with supplied ID: %v", err),
		)
	}
	return &empty.Empty{}, nil
}

func (s *PersonStore) DeleteRequest(_ context.Context, req *example.DeletePersonRequest) (*empty.Empty, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	_, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete blog with id %s: %v", req.GetId(), err))
	}
	return &empty.Empty{}, nil
}

func NewPersonStore(db *mongo.Database) *PersonStore {
	return &PersonStore{collection: db.Collection("persons")}
}
