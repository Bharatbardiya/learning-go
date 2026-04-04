package main

import (
	"embed"
	"fmt"
	"log"
)

////go:embed hello.txt
//var data string

//go:embed public
var public embed.FS

func main() {

	data, err := public.ReadFile("public/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
