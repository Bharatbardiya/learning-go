package main

import (
	"fmt"
)

func main() {
	message := make(chan string, 3)

	fmt.Println("sending message to Buffered channel")
	message <- "first message"
	message <- "second message"
	message <- "third message"
	// message <- "fourth message"

	fmt.Println("reciving message from Buffered channel")
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
