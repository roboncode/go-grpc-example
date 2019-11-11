package main

import (
	"aaa/api"
	aaa "aaa/generated"
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	// :: Connect :: //
	client, err := api.Connect()
	if err != nil {
		panic(err)
	}

	// :: Create :: //
	result, err := client.CreatePerson(context.Background(), &aaa.Person{
		Name: "My Name",
	})
	if err != nil {
		panic(err)
	}
	id := result.Id
	fmt.Println("")
	color.HiMagenta("Create")
	color.HiMagenta("----------------------------")
	fmt.Println(id)

	// :: Get :: //
	item, err := client.GetPerson(context.Background(), &aaa.Person_Id{Id: id})
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	color.HiMagenta("Get")
	color.HiMagenta("----------------------------")
	var bItem, _ = json.MarshalIndent(item, "", "   ")
	fmt.Println(string(bItem))

	// :: List :: //
	listResult, err := client.GetPersons(context.Background(), &aaa.Person_Filters{
		Enabled: false,
		Type:    0,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	color.HiMagenta("List")
	color.HiMagenta("----------------------------")
	var bList, _ = json.MarshalIndent(listResult, "", "   ")
	fmt.Println(string(bList))

	// :: Update :: //
	item.Name = "Name Override"
	updateResult, err := client.UpdatePerson(context.Background(), item)
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	color.HiMagenta("Update")
	color.HiMagenta("----------------------------")
	var bUpdated, _ = json.MarshalIndent(updateResult, "", "   ")
	fmt.Println(string(bUpdated))

	// :: Delete :: //
	_, err = client.DeletePerson(context.Background(), &aaa.Person_Id{Id: id})
	if err != nil {
		panic(err)
	}
	fmt.Println("")
	color.HiMagenta("Delete")
	color.HiMagenta("----------------------------")
	fmt.Println("ok")
	fmt.Println("")
}
