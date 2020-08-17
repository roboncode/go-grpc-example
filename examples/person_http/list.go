package person_http

import (
	"example/api"
	"example/examples/person_http/models"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type Persons struct {
	Items []models.Person `json:"items"`
}

func List(client *resty.Client) {
	logrus.Println("")
	logrus.Println("Get Persons")
	logrus.Println("----------------------------")

	var persons Persons
	res, err := client.R().
		EnableTrace().
		SetResult(&persons).
		Get("http://" + api.HttpAddress() + "/api/persons")
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Println(res)
}
