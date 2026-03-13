package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var name any

	fmt.Println(name)
	fmt.Println(1 + 1)
	fmt.Println(3.14)
	fmt.Println(true, false)
	fmt.Printf("%+v\n", []int{1, 2, 3})

	var greeting string
	greeting = "Hello World!!!"
	fmt.Println(greeting)

	var num int
	num = 10
	fmt.Println(num)

	var isNoogler bool
	isNoogler = true
	fmt.Println(isNoogler)

	var firstName, lastName string
	firstName = "Bharat"
	lastName = "Bardiya"

	fmt.Println(firstName, lastName)

	leetcodeId := "bharat_bardiya"
	fmt.Println(leetcodeId)

	var year = 2026
	fmt.Println(year)
}
