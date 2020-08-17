package person_http

import "github.com/go-resty/resty/v2"

func Connect() *resty.Client {
	client := resty.New()
	return client
}
