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
	pkg.PersonStoreServer
}

func (s *PersonStore) Create(ctx context.Context, req *pkg.Person) (*pkg.Person_Id, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
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
	return &pkg.Person_Id{
		Id: result.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}

func (s *PersonStore) Get(ctx context.Context, req *pkg.Person_Id) (*pkg.Person, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
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

	response := &pkg.Person{
		Id:        oid.Hex(),
		Name:      data.Name,
		Type:      data.Type,
		Enabled:   data.Enabled,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return response, nil
}

func (s *PersonStore) List(ctx context.Context, req *pkg.Person_Filters) (*pkg.Persons, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := s.collection.Find(context.Background(), bson.M{
		"type":    req.Type,
		"enabled": req.Enabled,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())

	var items = make([]*pkg.Person, 0)
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

		items = append(items, &pkg.Person{
			Id:        data.Id.Hex(),
			Name:      data.Name,
			Type:      data.Type,
			Enabled:   data.Enabled,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return &pkg.Persons{Items: items}, nil
}

func (s *PersonStore) Update(ctx context.Context, req *pkg.Person) (*pkg.Person, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
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
	result := s.collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'data'
	data := Person{}
	err = result.Decode(&data)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find Person with supplied ID: %v", err),
		)
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
	return &pkg.Person{
		Id:        oid.Hex(),
		Name:      data.Name,
		Type:      data.Type,
		Enabled:   data.Enabled,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (s *PersonStore) Delete(ctx context.Context, req *pkg.Person_Id) (*empty.Empty, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete blog with id %s: %v", req.GetId(), err))
	}
	return &empty.Empty{}, nil
}

func NewPersonStore(db *mongo.Database) *PersonStore {
	return &PersonStore{collection: db.Collection("persons")}
}
