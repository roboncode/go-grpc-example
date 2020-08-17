package person_grpc

import (
	"context"
	example "example/generated"
	"github.com/sirupsen/logrus"
)

func Create(client example.PersonServiceClient) string {
	logrus.Println("")
	logrus.Println("Create Person")
	logrus.Println("----------------------------")

	res, err := client.CreatePerson(context.Background(), &example.CreatePersonRequest{
		Name: "My Name",
	})
	if err != nil {
		logrus.Fatalln(err)
	}
	id := res.Id

	logrus.Println(id)

	return id
}
