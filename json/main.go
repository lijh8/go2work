package main // main executable package
// package json // non executable package

// a package in a module can be with different name as the module name;
// a package in a module can be main executable or non executable package;
// there can be only one package in module root directory;
// there can be sub module or sub package in its sub directory;

import (
	"encoding/json"
	"log"
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

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	var LOG = log.Println
	// var LOGF = log.Printf

	filename := "data.json"
	infile, err := os.Open(filename)
	if err != nil {
		LOG(err)
		return
	}
	defer infile.Close()

	var person Person

	decoder := json.NewDecoder(infile)
	if err := decoder.Decode(&person); err != nil {
		LOG(err)
		return
	}

	person.Age = 35
	person.Address.PostalCode = "12345"
	person.Hobbies[0] = "reading"

	outFile, err := os.Create(filename)
	if err != nil {
		LOG(err)
		return
	}
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(person); err != nil {
		LOG(err)
		return
	}
}
