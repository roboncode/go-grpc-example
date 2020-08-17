package main

import (
	"example/generated"
	"example/internal/connectors"
	"example/internal/grpc"
	"example/internal/healthcheck"
	"example/internal/http"
	"example/internal/service"
	"example/internal/store"
	"github.com/sirupsen/logrus"
	googleGrpc "google.golang.org/grpc"
)

const ServiceName = "example"

func connectToMongo() connectors.MongoConnector {
	mongoConnection := connectors.NewMongoConnection()
	if err := mongoConnection.Init(); err != nil {
		logrus.Fatalln(err)
	}
	return mongoConnection
}

func setupStores(conn connectors.MongoConnector) store.Store {
	s := store.NewStore()
	s.Set(store.PersonStoreName, store.NewPersonStore(conn.GetDatabase()))
	return s
}

func startGrpcServer(shutdown <-chan bool, stores store.Store) grpc.Server {
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
			defer logrus.Fatalln(err)
			<-shutdown
		}
	}()
	return grpcServer
}

func startHealthCheckServer(shutdown <-chan bool, grpcServer grpc.Server) healthcheck.Server {
	healthCheckServer := healthcheck.NewServer(ServiceName)
	go func() {
		if err := healthCheckServer.Serve(grpcServer.Instance()); err != nil {
			logrus.Fatalln(err)
			<-shutdown
		}
	}()
	return healthCheckServer
}

func startHttpServer(shutdown <-chan bool) http.Server {
	httpServer := http.NewServer()
	go func() {
		if err := httpServer.Serve(); err != nil {
			logrus.Fatalln(err)
			<-shutdown
		}
	}()
	return httpServer
}

func main() {
	shutdown := make(chan bool)

	conn := connectToMongo()
	stores := setupStores(conn)
	grpcServer := startGrpcServer(shutdown, stores)
	startHealthCheckServer(shutdown, grpcServer)
	startHttpServer(shutdown)

	<-shutdown
}
