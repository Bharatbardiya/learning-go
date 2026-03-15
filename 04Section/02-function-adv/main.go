package main

import (
	"fmt"
)

func factorial(n int) int{
	if n==0 || n==1 {
		return 1;
	}
	return n * factorial(n-1);
}

func intCounter() func() int {
	i:= 0;

	return func() int {
		i++;
		return i;
	}
}

func main() {

	fmt.Printf("Factorail of %d : %d\n", 4, factorial(4));

	nextInt := intCounter();

	fmt.Println(nextInt());
	fmt.Println(nextInt());
	fmt.Println(nextInt());
	fmt.Println(nextInt());
	

}