package main

import "fmt"


func main() {

	var arr [2]int;
	fmt.Println(arr);
	arr[0] = 100;
	arr[1] = 101;


	primes := [...]int {2,4,5,7,11};
	for i:= 0; i<len(primes); i++ {
		fmt.Printf("%d,", primes[i]);
	}
	fmt.Printf("\n");


	var matrix = [2][3]int{{1,2,3},{5,6,7}};

	for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix[0]); j++ {
			fmt.Printf("%d, ", matrix[i][j]);
		}
		fmt.Printf("\n");
	}
}