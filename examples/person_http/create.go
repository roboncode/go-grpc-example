package person_http

import (
	"encoding/json"
	"example/api"
	"example/examples/person_http/models"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func Create(client *resty.Client) string {
	logrus.Println("")
	logrus.Println("Create Person")
	logrus.Println("----------------------------")

	var person models.Person

	body := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  "Rob Taylor",
		Email: "roboncode@gmail.com",
	}
	b, _ := json.Marshal(body)

	_, err := client.R().
		EnableTrace().
		SetBody(string(b)).
		SetResult(&person).
		Post("http://" + api.HttpAddress() + "/api/persons")
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Println(person.Id)

	return person.Id
}
