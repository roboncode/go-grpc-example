package person_http

import (
	"encoding/json"
	"example/api"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
)

func Update(client *resty.Client, id string) {
	logrus.Println("")
	logrus.Println("Update Person")
	logrus.Println("----------------------------")

	body := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  "Name Override",
		Email: "override@gmail.com",
	}
	b, _ := json.Marshal(body)

	res, err := client.R().
		EnableTrace().
		SetBody(string(b)).
		SetResult(&empty.Empty{}).
		Get("http://" + api.HttpAddress() + fmt.Sprintf("/api/persons/%s", id))
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Println(res)
}
