package main

import (
	"fmt"
	"strings"
)

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

func (e MathError) Error() string {
	var inputs []string
	if e.Operation == "Devision" {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}

	return fmt.Sprintf("Math error is %s, (%s): %s",
		e.Operation,
		strings.Join(inputs, ", "),
		e.Message)
}

func Sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, &MathError{
			Operation: "Devision",
			InputA:    a,
			InputB:    b,
			Message:   "Devision by zero is not allowed",
		}
	}
	return a / b, nil
}

func main() {
	fmt.Println("hello world")
	sum := Sum(2, 3, 4, 4)
	fmt.Println(sum)
	division, err := Division(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(division)
	}
}
