package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels in go lang")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// Read only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {

		for val := range myChannel {
			fmt.Println(val)
		}
		wg.Done()
	}(myChannel, wg)

	//Write only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// myChannel <- 0
		// myChannel <- 5
		close(myChannel)
		// myChannel <- 4
		wg.Done()
	}(myChannel, wg)

	wg.Wait()

}
