package main

import (
	"aaa/api"
	"aaa/pkg"
	"aaa/tools/log"
	"context"
	"encoding/json"
)

func main() {
	// :: Connect :: //
	client, err := api.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	// :: Create :: //
	result, err := client.CreatePerson(context.Background(), &pkg.Person{
		Name: "My Name",
	})
	if err != nil {
		log.Fatalln(err)
	}
	id := result.Id
	log.Println("")
	log.Println("Create")
	log.Println("----------------------------")
	log.Println(id)

	// :: Get :: //
	item, err := client.GetPerson(context.Background(), &pkg.Person_Id{Id: id})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("")
	log.Println("Get")
	log.Println("----------------------------")
	var bItem, _ = json.MarshalIndent(item, "", "   ")
	log.Println(string(bItem))

	// :: List :: //
	listResult, err := client.GetPersons(context.Background(), &pkg.Person_Filters{
		Enabled: false,
		Type:    0,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("")
	log.Println("List")
	log.Println("----------------------------")
	var bList, _ = json.MarshalIndent(listResult, "", "   ")
	log.Println(string(bList))

	// :: Update :: //
	item.Name = "Name Override"
	updateResult, err := client.UpdatePerson(context.Background(), item)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("")
	log.Println("Update")
	log.Println("----------------------------")
	var bUpdated, _ = json.MarshalIndent(updateResult, "", "   ")
	log.Println(string(bUpdated))

	// :: Delete :: //
	_, err = client.DeletePerson(context.Background(), &pkg.Person_Id{Id: id})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("")
	log.Println("Delete")
	log.Println("----------------------------")
	log.Println("ok")
	log.Println("")
}
