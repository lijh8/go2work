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
	filename := "data.json"
	infile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer infile.Close()

	var person Person

	decoder := json.NewDecoder(infile)
	if err := decoder.Decode(&person); err != nil {
		fmt.Println(err)
		return
	}

	person.Age = 35
	person.Address.PostalCode = "12345"
	person.Hobbies[0] = "reading"

	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(person); err != nil {
		fmt.Println(err)
		return
	}
}
