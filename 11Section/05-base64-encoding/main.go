package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	data := "Welcome to Base64 Encoding"

	enc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(enc)

	fmt.Println([]byte(data))
	base64Str := "V2VsY29tZSB0byBCYXNlNjQgRW5jb2Rpbmc="

	dec, err := base64.StdEncoding.DecodeString(base64Str)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dec))
}
