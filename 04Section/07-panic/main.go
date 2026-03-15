package main

import "fmt"

func mightPanic(flag bool) {
	if flag {
		panic("something bad happened")
	}
	fmt.Println("returning without panic!!")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic", r)
		}
	}()
	mightPanic(true)
}
func main() {

	recoverable()

}
