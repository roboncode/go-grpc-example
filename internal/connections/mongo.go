package connections

import (
	"context"
	"example/util/env"
	"example/util/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	MongoAddress     = env.Var("MONGO_ADDR").Default("mongodb://localhost:27017").Desc("mongo address").String()
	MongoPingTimeout = env.Var("MONGO_PING_TIMEOUT").Default(60).Desc("mongo ping timeout").Duration()
	MongoDatabase    = env.Var("MONGO_DATABASE").Default("default").Desc("mongo database").String()
)

type MongoConnection interface {
	Init() error
	Connect(address string, pingTimeout time.Duration) (*mongo.Client, error)
	GetDatabase() *mongo.Database
}

type mongoConnection struct {
	client   *mongo.Client
	database *mongo.Database
}

func (m *mongoConnection) GetDatabase() *mongo.Database {
	return m.database
}

func (m *mongoConnection) Init() error {
	mongoClient, err := m.Connect(MongoAddress, MongoPingTimeout)
	if err != nil {
		return err
	}
	m.client = mongoClient
	m.database = mongoClient.Database(MongoDatabase)
	return nil
}

func (m *mongoConnection) Connect(address string, pingTimeout time.Duration) (*mongo.Client, error) {
	log.Infoln("connecting to mongo")
	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		log.Fatalln("failed to connect to mongo", err)
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), pingTimeout*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln("mongo health check failed", err)
		return nil, err
	}

	log.Infoln("connected to mongo")

	return client, nil
}

func NewMongoConnection() MongoConnection {
	return &mongoConnection{}
}
