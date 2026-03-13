package main

import (
	"fmt"
)


func main () {

	var friends = []string {"Jhon", "Alise", "ramesh"}
	fmt.Printf("friends : %+v\n", friends);

	arr :=  make([]int, 3, 5);

	fmt.Printf("arr: %+v, size: %d, cap %d\n", arr, len(arr), cap(arr));

	arr = append(arr, 1);
	arr = append(arr, 2);
	arr = append(arr, 3);

	fmt.Printf("arr: %+v, size: %d, cap %d\n", arr, len(arr), cap(arr));

	fmt.Println("new element in array:", arr[3:6]);

}