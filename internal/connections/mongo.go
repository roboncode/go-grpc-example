package connections

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func ConnectToMongo(address string, pingTimeout time.Duration) (*mongo.Client, error) {
	fmt.Print("Connecting to mongo... ")
	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		fmt.Println("error")
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), pingTimeout*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("error... ping failed")
		return nil, err
	}

	fmt.Println("success")

	return client, nil
}
