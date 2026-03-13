package main

import "fmt"

func main() {

	// c style loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while style loop
	counter := 3
	for counter > 0 {
		fmt.Println(counter)
		counter--
	}

	// python style loop
	languages := [4]string{"go", "python", "java", "c"}
	for index, value := range languages {
		fmt.Println(index, value)
	}

	for _, value := range languages {
		fmt.Println(value)
	}
}
