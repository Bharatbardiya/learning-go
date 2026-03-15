package main

import (
	"errors"
	"fmt"
    "strings"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("divide by zero not possible")
		return 0, err
	}
	return a / b, nil
}

func splitFullName(fullName string) (firstName, lastName string){
    parts := strings.Split(fullName, " ");
    firstName, lastName = parts[0], parts[1]
    return
}

func main() {

	ans, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}

    firstName, lastName := splitFullName("Bharat Bardiya");
    fmt.Printf("firstName : %s, lastName : %s\n", firstName, lastName);
}
