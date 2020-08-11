package main

import (
	"example/api"
	"example/internal/connections"
	"example/internal/grpc"
	"example/internal/server"
	"example/internal/store"
	"example/tools/log"
)

func main() {
	shutdown := make(chan bool)

	// Setup connections
	mongoConnection := connections.NewMongoConnection()
	if err := mongoConnection.Init(); err != nil {
		log.Fatalln(err)
	}

	// Setup store
	s := store.NewStore()
	s.Set(store.PersonStoreName, store.NewPersonStore(mongoConnection.Database))

	// Setup servers
	appServer := server.NewServer(s)
	grpcServer := grpc.NewServer()

	go func() {
		log.Infof("Listening to gRPC on %s\n", api.GrpcAddr)

		if err := grpcServer.Serve(appServer); err != nil {
			defer log.Fatalln(err)
			<-shutdown
		}
	}()

	<-shutdown
}
