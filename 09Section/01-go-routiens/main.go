package main

import (
	"fmt"
	"time"
)

func sayHello(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(message)
}

func main() {
	fmt.Println("Hello world from main go routien")

	go sayHello("Hello world 1", time.Second)
	go sayHello("Hello world 2", time.Second)
	go sayHello("Hello world after 2 Second", 2*time.Second)
	go sayHello("Hello world after 3 Second", 3*time.Second)

	fmt.Println("last message from main() go routien")
	time.Sleep(time.Second * 2)
}
