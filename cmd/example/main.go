package main

import (
	"example/generated"
	"example/internal/connectors"
	"example/internal/servers/grpc"
	"example/internal/servers/healthcheck"
	"example/internal/servers/http"
	"example/internal/services"
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

func setupStores(conn connectors.MongoConnector) {
	store.Instance().Person = store.NewPersonStore(conn.GetDatabase())
}

func startGrpcServer(shutdown <-chan bool) grpc.Server {
	var personService = services.NewPersonService()
	var httpService = services.NewHttpService(personService)

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
	setupStores(conn)
	grpcServer := startGrpcServer(shutdown)
	startHealthCheckServer(shutdown, grpcServer)
	startHttpServer(shutdown)

	<-shutdown
}
