package main

import (
	"fmt"
)

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			r, ok := <-jobs
			if ok {
				fmt.Println("recived job ", r)
			} else {
				done <- true
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sending...", i)
	}
	close(jobs)

	<-done
}
