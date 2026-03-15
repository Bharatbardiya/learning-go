package main

import "fmt"

type Person interface {
	GetName() string
	GetId() int
}

type Employee struct {
	id   int
	name string
}

func (emp *Employee) GetName() string {
	return emp.name
}

type BusinessMan struct {
	id       int
	fullName string
}

func (emp *BusinessMan) GetName() string {
	return emp.fullName
}

func main() {
	bharat := Employee{1, "Bharat"}
	fmt.Println(bharat.GetName())

	bharatBardiya := BusinessMan{1, "Bharat Bardiya"}
	fmt.Println(bharatBardiya.GetName())
}
