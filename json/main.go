package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Street     string `json:"street"`
	PostalCode string `json:"postalCode"`
}

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	City    string   `json:"city"`
	Address Address  `json:"address"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var person Person
	err = json.NewDecoder(file).Decode(&person)
	if err != nil {
		fmt.Println(err)
		return
	}

	person.Age = 35
	person.Address.PostalCode = "12345"
	person.Hobbies[0] = "reading"

	modifiedData, err := json.MarshalIndent(person, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("new.json", modifiedData, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Changes written to new.json")
}

/*
{
    "name": "Alice",
    "age": 31,
    "city": "Wonderland",
    "address": {
        "street": "123 Rabbit Hole",
        "postalCode": "67890"
    },
    "hobbies": [
        "writing",
        "adventuring",
        "tea parties"
    ]
}
*/
