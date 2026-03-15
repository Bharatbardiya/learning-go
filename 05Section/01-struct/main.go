package main

import (
	"fmt"
	"time"
)

type Employee struct {
	firstName, lastName, position string
	Salary, ID                    int
	isActive                      bool
	JoinedAt                      time.Time
}

func newEmployee(firstName, lastName, position string, ID int) Employee {
	return Employee{
		firstName: firstName,
		lastName:  lastName,
		position:  position,
		ID:        ID,
		isActive:  true,
	}
}

func (emp *Employee) fullName() string {
	return emp.firstName + " " + emp.lastName
}

func (emp *Employee) Activate() {
	emp.isActive = true
}

func (emp *Employee) Deactivate() {
	emp.isActive = false
}

func (emp *Employee) setJoningDate(t time.Time) {
	emp.JoinedAt = t
}

func main() {
	Bharat := Employee{
		firstName: "Bharat",
		lastName:  "Bardiya",
		position:  "SWE(L3)",
		Salary:    1000,
		ID:        1,
		JoinedAt:  time.Now(),
	}
	fmt.Printf("%+v\n", Bharat)
	fmt.Printf("%s\n", Bharat.fullName())
	Bharat.Activate()
	Bharat.setJoningDate(time.Now().Add(1000 * time.Minute))
	fmt.Printf("%+v\n", Bharat)

	Jhon := newEmployee("Jhon", "carner", "Manager-1", 2)
	Jhon.Salary = 2000

	jhonPtr := &Jhon
	fmt.Println(jhonPtr)
}
