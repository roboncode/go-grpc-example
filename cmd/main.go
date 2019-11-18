package main

import (
	"aaa/api"
	"aaa/internal/connections"
	"aaa/internal/grpc"
	"aaa/internal/server"
	"aaa/internal/store"
	"flag"
	"log"
)

func main() {
	shutdown := make(chan bool)

	flag.Parse()

	// Setup connections
	mongoConnection := connections.NewMongoConnection()
	if err := mongoConnection.Init(); err != nil {
		log.Fatalln(err)
	}

	// Setup store
	s := store.NewStore()
	s.Set(store.PersonStoreName, store.NewPersonStore(mongoConnection.Database))

	// Setup servers
	appServer := server.NewServer(&s)
	grpcServer := grpc.NewServer()
	go func() {
		log.Printf("Listening to gRPC on %s\n", api.GrpcAddr)

		if err := grpcServer.Serve(appServer); err != nil {
			defer log.Fatalln(err)
			<-shutdown
		}
	}()

	<-shutdown
}
