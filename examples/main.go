package main

import (
	"example/examples/person_grpc"
	"example/examples/person_http"
	"github.com/sirupsen/logrus"
)

func runGrpcExample() {
	logrus.Println("*** GRPC EXAMPLE ***")
	client := person_grpc.Connect()
	id := person_grpc.Create(client)
	person_grpc.Get(client, id)
	person_grpc.List(client)
	person_grpc.Update(client, id)
	person_grpc.Delete(client, id)
}

func runHttpExample() {
	logrus.Println("*** HTTP EXAMPLE ***")
	client := person_http.Connect()
	id := person_http.Create(client)
	person_http.Get(client, id)
	person_http.List(client)
	person_http.Update(client, id)
	person_http.Delete(client, id)
}

func main() {
	runGrpcExample()
	runHttpExample()
}
