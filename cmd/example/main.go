package main

import (
	"example/generated"
	"example/internal/connections"
	"example/internal/grpc"
	"example/internal/healthcheck"
	"example/internal/http"
	"example/internal/service"
	"example/internal/store"
	log "github.com/sirupsen/logrus"
	googleGrpc "google.golang.org/grpc"
)

const ServiceName = "example"

func connectToMongo() connections.MongoConnection {
	mongoConnection := connections.NewMongoConnection()
	if err := mongoConnection.Init(); err != nil {
		log.Fatalln(err)
	}
	return mongoConnection
}

func setupStores(conn connections.MongoConnection) store.Store {
	s := store.NewStore()
	s.Set(store.PersonStoreName, store.NewPersonStore(conn.GetDatabase()))
	return s
}

func setupGrpcServer(shutdown <-chan bool) grpc.Server {
	conn := connectToMongo()
	stores := setupStores(conn)

	var personService = service.NewPersonService(stores)
	var httpService = service.NewHttpService(personService)

	var opts = grpc.Options{
		ServiceRegistration: func(s *googleGrpc.Server) {
			example.RegisterPersonServiceServer(s, personService)
			example.RegisterHttpServiceServer(s, httpService)
		},
	}

	grpcServer := grpc.NewServer(&opts)
	go func() {
		if err := grpcServer.Serve(); err != nil {
			defer log.Fatalln(err)
			<-shutdown
		}
	}()
	return grpcServer
}

func setupHealthCheckServer(shutdown <-chan bool, grpcServer grpc.Server) healthcheck.Server {
	healthCheckServer := healthcheck.NewServer(ServiceName)
	go func() {
		if err := healthCheckServer.Serve(grpcServer.Instance()); err != nil {
			log.Fatalln(err)
			<-shutdown
		}
	}()
	return healthCheckServer
}

func setupHttpServer(shutdown <-chan bool) http.Server {
	httpServer := http.NewServer()
	go func() {
		if err := httpServer.Serve(); err != nil {
			log.Fatalln(err)
			<-shutdown
		}
	}()
	return httpServer
}

func main() {
	shutdown := make(chan bool)

	grpcServer := setupGrpcServer(shutdown)
	setupHealthCheckServer(shutdown, grpcServer)
	setupHttpServer(shutdown)

	<-shutdown
}
