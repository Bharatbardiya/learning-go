package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone_number"`
	IsActive bool   `json:"is_active" xml:"is_active"`
	Role     string `json:"-" xml:"-"`
}

var payload = `{"name":"alias", "age":20, "role":"admin"}`

func main() {

	var usr user
	// decoder require a source/reader

	dec := json.NewDecoder(strings.NewReader(payload))
	if err := dec.Decode(&usr); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("usr2 : %+v\n", usr)
}
