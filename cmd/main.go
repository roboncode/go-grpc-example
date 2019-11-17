package main

import (
	"aaa/api"
	"aaa/internal/connections"
	"aaa/internal/grpc"
	"aaa/internal/server"
	"aaa/internal/store"
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func main() {
	shutdown := make(chan bool)

	flag.Parse()

	// Setup drivers
	mongoDriver := connections.NewMongoConnection()
	mongoDriver.Init()

	// Setup store
	s := store.NewStore()
	s.Set(store.PersonStoreName, store.NewPersonStore(mongoDriver.Database))

	// Setup servers
	appServer := server.NewServer(&s)
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
