package person_grpc

import (
	"context"
	"encoding/json"
	"example/generated"
	"github.com/sirupsen/logrus"
)

func List(client example.PersonServiceClient) {
	logrus.Println("")
	logrus.Println("Get Persons")
	logrus.Println("----------------------------")

	persons, err := client.GetPersons(context.Background(), &example.GetPersonsRequest{})
	if err != nil {
		logrus.Fatalln(err)
	}
	var bList, _ = json.MarshalIndent(persons, "", "   ")

	logrus.Println(string(bList))
}
