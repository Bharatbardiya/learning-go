package main

import (
	"encoding/json"
	"log"
	"os"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone_number"`
	IsActive bool   `json:"is_active" xml:"is_active"`
	Role     string `json:"-" xml:"-"`
}

func main() {

	usr := user{
		Name:     "John Doe",
		Age:      21,
		IsActive: true,
		Role:     "User",
	}

	// encoder will need a destination where it can write the data
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(usr); err != nil {
		log.Fatal(err)
	}

}
