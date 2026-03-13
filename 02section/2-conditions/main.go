package main

import "fmt"

func main() {

	userAccess := map[string]bool{
		"Bharat": true,
		"Vinod":  false,
	}

	if hasAccess, ok := userAccess["Bharat"]; ok && hasAccess {
		fmt.Println("Bharat have access to the system")
	} else {
		fmt.Println("Bharat is not accessible to the system")
	}
}
