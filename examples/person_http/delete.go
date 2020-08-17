package person_http

import (
	"example/api"
	"example/generated"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func Delete(client *resty.Client, id string) {
	logrus.Println("")
	logrus.Println("Delete Person")
	logrus.Println("----------------------------")

	res, err := client.R().
		EnableTrace().
		SetResult(&example.Person{}).
		Delete("http://" + fmt.Sprintf(api.HttpAddress()+"/api/persons/%s", id))
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Println(res)
}
