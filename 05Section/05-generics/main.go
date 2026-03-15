package main

import "fmt"

type Number interface {
	int | float32 | float64
}

func Sum[T Number](numbers ...T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	total := Sum(1, 2, 3, 4)
	totalFloat := Sum(2.33, 4.54, 1.21)

	fmt.Printf("total is %d, type is %T\n", total, total)
	fmt.Printf("totalFloat is %f, type is %T\n", totalFloat, totalFloat)
}
