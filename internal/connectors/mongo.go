package connectors

import (
	"context"
	"example/util/env"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	MongoAddress     = env.Var("MONGO_ADDR").Default("mongodb://localhost:27017").Desc("mongo address").String()
	MongoPingTimeout = env.Var("MONGO_PING_TIMEOUT").Default(30).Min(1).Max(30).Desc("mongo ping timeout").Duration()
	MongoDatabase    = env.Var("MONGO_DATABASE").Default("default").Desc("mongo database").String()
)

type MongoConnector interface {
	Init() error
	Connect(address string, pingTimeout time.Duration) (*mongo.Client, error)
	GetDatabase() *mongo.Database
}

type mongoConnector struct {
	client   *mongo.Client
	database *mongo.Database
}

func (m *mongoConnector) GetDatabase() *mongo.Database {
	return m.database
}

func (m *mongoConnector) Init() error {
	mongoClient, err := m.Connect(MongoAddress, MongoPingTimeout)
	if err != nil {
		return err
	}
	m.client = mongoClient
	m.database = mongoClient.Database(MongoDatabase)
	return nil
}

func (m *mongoConnector) Connect(address string, pingTimeout time.Duration) (*mongo.Client, error) {
	log.Infoln("connecting to mongo")

	counter := 0
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				counter = counter + 1
				fmt.Print(".")
			}
		}
	}()

	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		done <- true
		fmt.Println()
		log.Fatalln("failed to connect to mongo", err)
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), pingTimeout*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		done <- true
		fmt.Println()
		log.Fatalln("mongo health check failed", err)
		return nil, err
	}

	done <- true
	fmt.Println()
	log.Infoln("connected to mongo")

	return client, nil
}

func NewMongoConnection() MongoConnector {
	return &mongoConnector{}
}
