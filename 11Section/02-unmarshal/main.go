package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string  `json:"name" xml:"name"`
	Age      int     `json:"age" xml:"age"`
	Phone    string  `json:"phone" xml:"phone_number"`
	IsActive bool    `json:"is_active" xml:"is_active"`
	Role     string  `json:"-" xml:"-"`
	Password string  `json:"-" xml:"-"`
	Profile  profile `json:"profile" xml:"profile"`
}

type profile struct {
	URL string `json:"url" xml:"url"`
}

var payload = `{
	"Name":"Bharat",
	"Age":21,
	"Phone":"123-1234-1223",
	"IsActive":true,
	"Password":"Dummy@123",
	"profile" : {"url": "https://leetcode.com/u/bharat_bardiya"}
}`

func main() {

	var usr user

	err := json.Unmarshal([]byte(payload), &usr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", usr)
	fmt.Printf("%+v\n", usr.Role == "")

	data, err := json.MarshalIndent(usr, "", "")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
