package store

import (
	"context"
	"example/internal/models"
	"example/internal/types/fieldtype"
	"example/internal/types/mongotype"
	"example/internal/types/objecttype"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PersonFilters struct {
	Status int32 `bson:"status,omitempty"`
	Skip   int64 `bson:"skip,omitempty"`
	Limit  int64 `bson:"offset,omitempty"`
}

type PersonStore interface {
	CreatePerson(person *models.Person) error
	GetPerson(id string) (*models.Person, error)
	GetPersons(filters *PersonFilters) ([]models.Person, error)
	UpdatePerson(id string, person *models.Person) error
	DeleteRequest(id string) error
}

type personStore struct {
	collection *mongo.Collection
}

func (s *personStore) CreatePerson(person *models.Person) error {
	var now = time.Now()
	person.CreatedAt = &now
	person.UpdatedAt = &now
	result, err := s.collection.InsertOne(context.Background(), &person)
	if err != nil {
		return err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	person.Id = &oid
	return nil
}

func (s *personStore) GetPerson(id string) (*models.Person, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := s.collection.FindOne(context.Background(), bson.M{fieldtype.Id: oid})
	var person models.Person
	if err := result.Decode(&person); err != nil {
		//return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find Person with Object Id %s: %v", id), err)
		return nil, err
	}
	return &person, nil
}

func (s *personStore) GetPersons(personFilters *PersonFilters) ([]models.Person, error) {
	query := bson.M{}
	if personFilters.Status > 0 {
		query[fieldtype.Status] = personFilters.Status
	}

	findOpts := options.FindOptions{}
	findOpts.SetSkip(personFilters.Skip)
	findOpts.SetLimit(personFilters.Limit)

	cursor, err := s.collection.Find(context.Background(), query, &findOpts)
	if err != nil {
		return nil, err
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())

	var persons = make([]models.Person, 0)
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		var person models.Person
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(&person)
		// check error
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}

func (s *personStore) UpdatePerson(id string, person *models.Person) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := s.collection.FindOneAndUpdate(context.Background(), filter, bson.M{mongotype.SET: person}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'data'
	data := models.Person{}
	err = result.Decode(&data)
	if err != nil {
		return err
	}
	return nil
}

func (s *personStore) DeleteRequest(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.collection.DeleteOne(context.Background(), bson.M{fieldtype.Id: oid})
	if err != nil {
		return err
	}
	return nil
}

func NewPersonStore(db *mongo.Database) PersonStore {
	return &personStore{collection: db.Collection(objecttype.Persons)}
}
