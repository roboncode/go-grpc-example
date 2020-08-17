package person_grpc

import (
	"example/api"
	"example/generated"
	"github.com/sirupsen/logrus"
)

func Connect() example.PersonServiceClient {
	conn, err := api.Connect(api.GrpcAddress())
	if err != nil {
		logrus.Fatalln(err)
	}
	return example.NewPersonServiceClient(conn)
}
