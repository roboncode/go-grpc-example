package main

import (
	"example/api"
	example "example/generated"
	"example/internal/connections"
	"example/internal/grpc"
	"example/internal/healthcheck"
	"example/internal/http"
	"example/internal/service"
	"example/internal/store"
	"example/util/log"
)

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

func setupAppServer(s store.Store) example.AppServiceServer {
	return service.NewServer(s)
}

func setupGrpcServer(shutdown <-chan bool, appServiceServer example.AppServiceServer) grpc.Server {
	log.Infof("Listening to gRPC on %s\n", api.GrpcAddress())
	grpcServer := grpc.NewServer()
	go func() {
		if err := grpcServer.Serve(appServiceServer); err != nil {
			defer log.Fatalln(err)
			<-shutdown
		}
	}()
	return grpcServer
}

func setupHealthCheckServer(shutdown <-chan bool, grpcServer grpc.Server) healthcheck.Server {
	log.Infoln("Listening to HealthCheck")
	healthCheckServer := healthcheck.NewServer()
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
	log.Infof("Listening to HTTP on %s\n", api.HttpAddress())
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

	conn := connectToMongo()
	stores := setupStores(conn)
	appServer := setupAppServer(stores)
	grpcServer := setupGrpcServer(shutdown, appServer)
	setupHealthCheckServer(shutdown, grpcServer)
	setupHttpServer(shutdown)

	<-shutdown
}
