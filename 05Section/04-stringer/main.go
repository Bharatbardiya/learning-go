package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (emp Employee) String() string {
	return fmt.Sprintf("Employee[id: %d, name: %s]", emp.id, emp.name)
}

type ID int

func (id ID) String() string {
	return fmt.Sprintf("custom Stringer : ID[%d]", id)
}
func main() {

	Bharat := Employee{1, "Bharat"}
	fmt.Println(Bharat)

	var myId ID
	myId = 30
	fmt.Println(myId)
}
