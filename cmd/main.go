package main

import (
	"aaa/api"
	"aaa/cmd/setup"
	"aaa/internal/grpc"
	"aaa/internal/server"
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func main() {
	shutdown := make(chan bool)

	flag.Parse()
	setup.Mongo()

	appServer := server.NewServer()
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
