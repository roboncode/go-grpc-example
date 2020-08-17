package person_grpc

import (
	"context"
	"encoding/json"
	"example/generated"
	"github.com/sirupsen/logrus"
)

func Update(client example.PersonServiceClient, id string) {
	logrus.Println("")
	logrus.Println("Update Person")
	logrus.Println("----------------------------")

	res, err := client.UpdatePerson(context.Background(), &example.UpdatePersonRequest{
		Id:     id,
		Name:   "Name Override",
		Status: example.Status_ACTIVE,
		Email:  "override@gmail.com",
	})
	if err != nil {
		logrus.Fatalln(err)
	}

	var bUpdated, _ = json.MarshalIndent(res, "", "   ")
	logrus.Println(string(bUpdated))
}
