package person_grpc

import (
	"context"
	"encoding/json"
	"example/generated"
	"github.com/sirupsen/logrus"
)

func Get(client example.PersonServiceClient, id string) {
	logrus.Println("")
	logrus.Println("Get Person")
	logrus.Println("----------------------------")

	item, err := client.GetPerson(context.Background(), &example.GetPersonRequest{Id: id})
	if err != nil {
		logrus.Fatalln(err)
	}

	var bItem, _ = json.MarshalIndent(item, "", "   ")
	logrus.Println(string(bItem))
}
