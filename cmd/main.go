package main

import (
	"aaa/api"
	"aaa/internal/connections"
	"aaa/internal/grpc"
	"aaa/internal/server"
	"aaa/tools/env"
	"flag"
	"fmt"
	"github.com/golang/glog"
)

var (
	mongoAddress     = env.String("AAA_MONGO_ADDR", "mongodb://localhost:27017", "mongo address")
	mongoPingTimeout = env.Duration("AAA_MONGO_PING_TIMEOUT", 2, "mongo ping timeout")
	mongoDatabase    = env.String("AAA_MONGO_DATABASE", "default", "mongo database")
)

func main() {
	shutdown := make(chan bool)

	flag.Parse()

	mongoClient, err := connections.ConnectToMongo(mongoAddress, mongoPingTimeout)
	if err != nil {
		panic(err)
	}

	appServer := server.NewServer()
	appServer.InitMongoStore(mongoClient.Database(mongoDatabase))

	grpcServer := grpc.NewServer()
	go func() {
		fmt.Printf("Listening to gRPC on %s\n", api.GrpcAddr)

		if err := grpcServer.Serve(appServer); err != nil {
			glog.Fatalln(err)
			<-shutdown
		}
	}()

	<-shutdown
}
