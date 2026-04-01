package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
}

func main() {
	message := make(chan string)

	go func() {
		message <- "sending message to string channel."
	}()

	go func() {
		message <- "sending message to string channel another time."
	}()

	usr := make(chan user)

	go func() {
		usr <- user{
			name: "Bharat Bardiya",
		}
	}()

	time.Sleep(time.Second)

	fmt.Println("about to get message from string channel")
	msg := <-message
	fmt.Println(msg)
	msg = <-message
	fmt.Println(msg)

	fmt.Println("about to get message from usr channel")
	u := <-usr
	fmt.Println(u)
}
