package main

import (
	"fmt"
)

func main(){

	// initializing with values
	studentGrade := map[string]int {
		"Alias" : 90,
		"Jhon" : 85,
		"Bharat" : 80,
	}

	fmt.Println(studentGrade);


	// updating value

	studentGrade["Bharat"] = 95;

	fmt.Printf("studentGrade updated ; %+v\n", studentGrade);

	key := "Jhon";
	if marks, ok := studentGrade[key]; ok {
		fmt.Printf("jhon's marks : %d\n", marks);
	}

	var config1 map[string]int;

	fmt.Printf("Config1 is nil : %+v\n", config1==nil);

	config2 := map[string]int{};
	fmt.Printf("Config2 is nil : %+v\n", config2==nil);

}