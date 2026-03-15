package main

import ("fmt")


func sum(numbers ...int) int{

	var sum int = 0;

	for _, number := range numbers{
		sum+= number;
	}
	return sum;
}

func config(numbers ...int){
	if len(numbers) > 0 {
		fmt.Println("Do something spacial");
	} else {
		fmt.Println("Default behaviour");
	}
}

func main(){

	fmt.Println("sum of numbers: ", sum(1,2,3));
	fmt.Println("sum of numbers: ", sum(1,2,3,4,5));

	config();
	config(1,2,3);

}