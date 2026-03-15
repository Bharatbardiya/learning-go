package main

import "fmt"

func main() {

	var arr [2]int
	fmt.Println(arr)
	arr[0] = 100
	arr[1] = 101

	primes := [...]int{2, 4, 5, 7, 11}
	for i := 0; i < len(primes); i++ {
		fmt.Printf("%d,", primes[i])
	}
	fmt.Printf("\n")
}
