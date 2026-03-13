package main

import (
	"fmt"
	"time"
)

func main() {

	day := "Monday"
	fmt.Println("Today is ", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Its Weekend!!, No work")

	case "Monday":
		fmt.Println("work day, lots of meeting!!")

	default:
		fmt.Println("mid-week")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	checkType := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Printf("int %d\n", t)
		case string:
			fmt.Printf("string %s\n", t)
		case bool:
			fmt.Printf("bool %t\n", t)
		default:
			fmt.Printf("unknown type %T\n", t)
		}
	}

	checkType(10)
	checkType("bharat")
	checkType(true)
	checkType(23.34)

}
