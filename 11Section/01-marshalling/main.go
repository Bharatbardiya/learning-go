package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age"	xml:"age"`
	Phone    string `json:"phone" xml:"phone_number"`
	IsActive bool   `json:"is_active" xml:"is_active"`
}

func main() {

	bharat := user{
		Name:     "Bharat",
		Age:      21,
		Phone:    "123-1234-1223",
		IsActive: true,
	}

	jsonData, err := json.Marshal(bharat)
	if err != nil {
		log.Fatal(err)
	}

	xmlData, err := xml.Marshal(bharat)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", bharat)
	fmt.Println(string(jsonData))

	fmt.Println(string(xmlData))
}
