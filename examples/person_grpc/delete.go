package person_grpc

import (
	"context"
	example "example/generated"
	"github.com/sirupsen/logrus"
)

func Delete(client example.PersonServiceClient, id string) {
	logrus.Println("")
	logrus.Println("Delete Person")
	logrus.Println("----------------------------")

	_, err := client.DeletePerson(context.Background(), &example.DeletePersonRequest{Id: id})
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Println("ok")
	logrus.Println("")
}
