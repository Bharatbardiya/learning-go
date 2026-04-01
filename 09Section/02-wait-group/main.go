package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println(message)
}

func JobsRunner(wg *sync.WaitGroup) {

	totalJobs := 5

	for i := 0; i < totalJobs; i++ {
		wg.Add(1)
		go sayHello(fmt.Sprintf("running Job %d", i+1), (time.Duration)(i+1)*time.Second, wg)
	}
}

func main() {

	var jobWg sync.WaitGroup

	JobsRunner(&jobWg)
	defer jobWg.Wait()

	fmt.Println("Hello world from main go routien")

	var wg sync.WaitGroup

	wg.Add(4)

	go sayHello("Hello world 1", time.Second, &wg)
	go sayHello("Hello world 2", time.Second, &wg)
	go sayHello("Hello world after 2 Second", 2*time.Second, &wg)
	go sayHello("Hello world after 3 Second", 3*time.Second, &wg)

	fmt.Println("last message from main() go routien")
	wg.Wait()

}
