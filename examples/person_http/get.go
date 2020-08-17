package person_http

import (
	"example/api"
	"example/examples/person_http/models"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func Get(client *resty.Client, id string) {
	logrus.Println("")
	logrus.Println("Get Person")
	logrus.Println("----------------------------")

	var person models.Person
	res, err := client.R().
		EnableTrace().
		SetResult(&person).
		Get("http://" + api.HttpAddress() + fmt.Sprintf("/api/persons/%s", id))
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Println(res)
}
