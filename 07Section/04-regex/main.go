package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	text := "this is a Go course"

	regGo, err := regexp.Compile(`Go`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(regGo.FindString(text))

	productList := "all products : X123, P122, Y23, P98a23"

	productRegex := regexp.MustCompile(`P\d+`)

	firstP := productRegex.FindString(productList)
	allP := productRegex.FindAllString(productList, -1)

	fmt.Println(firstP)
	fmt.Println(allP)
}
