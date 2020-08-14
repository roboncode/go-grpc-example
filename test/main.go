package main

import (
	"context"
	"encoding/json"
	"example/api"
	"example/generated"
	log "github.com/sirupsen/logrus"
)

func main() {
	// :: Connect To Client :: //
	conn, err := api.Connect(api.GrpcAddress())
	if err != nil {
		log.Fatalln(err)
	}
	client := example.NewAppServiceClient(conn)

	// :: Create :: //
	result, err := client.CreatePerson(context.Background(), &example.CreatePersonRequest{
		Name: "My Name",
	})
	if err != nil {
		log.Fatalln(err)
	}
	id := result.Id
	log.Println("")
	log.Println("Create Person")
	log.Println("----------------------------")
	log.Println(id)

	// :: Get :: //
	item, err := client.GetPerson(context.Background(), &example.GetPersonRequest{Id: id})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("")
	log.Println("Get Person")
	log.Println("----------------------------")
	var bItem, _ = json.MarshalIndent(item, "", "   ")
	log.Println(string(bItem))

	// :: List :: //
	persons, err := client.GetPersons(context.Background(), &example.GetPersonsRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("")
	log.Println("Get Persons")
	log.Println("----------------------------")
	var bList, _ = json.MarshalIndent(persons, "", "   ")
	log.Println(string(bList))

	// :: Update :: //
	updateResult, err := client.UpdatePerson(context.Background(), &example.UpdatePersonRequest{
		Id:     id,
		Name:   "Name Override",
		Status: example.Status_ACTIVE,
		Email:  "override@gmail.com",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("")
	log.Println("Update Person")
	log.Println("----------------------------")
	var bUpdated, _ = json.MarshalIndent(updateResult, "", "   ")
	log.Println(string(bUpdated))

	// :: Delete :: //
	_, err = client.DeletePerson(context.Background(), &example.DeletePersonRequest{Id: id})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("")
	log.Println("Delete Person")
	log.Println("----------------------------")
	log.Println("ok")
	log.Println("")
}
