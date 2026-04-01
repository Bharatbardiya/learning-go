package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func openConnection(done chan bool) {
	fmt.Println("Attempting connection....")

	if rand.Intn(100) > 50 {
		time.Sleep(2 * time.Second)
		fmt.Println("Connection succeeded")
	} else {
		fmt.Println("OOPS!!  Hanging connection!!")
		time.Sleep(1000 * time.Hour)
	}
	done <- true
}

func openConnectionWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	done := make(chan bool)
	go openConnection(done)

	select {
	case <-done:
		fmt.Println("Connection Successful")
	case <-ctx.Done():
		fmt.Println("Connection timed out")
	}
}

func main() {
	openConnectionWithTimeout()
}
