package connections

import (
	"aaa/tools/env"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var (
	MongoAddress     = env.String("AAA_MONGO_ADDR", "mongodb://localhost:27017", "mongo address")
	MongoPingTimeout = env.Duration("AAA_MONGO_PING_TIMEOUT", 2, "mongo ping timeout")
	MongoDatabase    = env.String("AAA_MONGO_DATABASE", "default", "mongo database")
)

type MongoConnection struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func (m *MongoConnection) Init() error {
	mongoClient, err := m.Connect(MongoAddress, MongoPingTimeout)
	if err != nil {
		return err
	}
	m.Client = mongoClient
	m.Database = mongoClient.Database(MongoDatabase)
	return nil
}

func (m *MongoConnection) Connect(address string, pingTimeout time.Duration) (*mongo.Client, error) {
	log.Print("Connecting to mongo")
	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		log.Fatalln("Error connecting to mongo")
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), pingTimeout*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln("Error pinging mongo")
		return nil, err
	}

	log.Println("Connected to mongo")

	return client, nil
}

func NewMongoConnection() *MongoConnection {
	return &MongoConnection{}
}
