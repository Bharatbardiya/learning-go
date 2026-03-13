// no earthmatic operations are allowed on pointers in go

package main

import ("fmt")

func modifyValue(numPtr *int) {

	*numPtr = *numPtr * 10;
	fmt.Printf("printing derefrence value : %d\n", *numPtr);
}

func main() {

	num := 10; 
	modifyValue(&num);

	fmt.Printf("num : %d\n", num);


	grade := 50;
	gradePtr := &grade;

	fmt.Printf("grade variable address : %v\n", gradePtr);
	fmt.Printf("grade variable address : %v\n", *(&gradePtr));
}