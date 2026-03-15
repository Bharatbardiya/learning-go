package main

import (
	"fmt"
	"os"
)

func simpleDefer() {
	fmt.Println("simpleDefer()")
	defer fmt.Println("deferred print simpleDefer()")
	fmt.Println("returning from simpleDefer()")
}

func lifoSimpleDefer() {
	fmt.Println("lifoDefer()")
	defer fmt.Println("first deferred print lifoDefer()")
	defer fmt.Println("second deferred print lifoDefer()")
	fmt.Println("returning from lifoDefer()")
}

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	defer fmt.Println("deferred print main()")

	simpleDefer()
	lifoSimpleDefer()

	fmt.Println("returning from main()")
}
