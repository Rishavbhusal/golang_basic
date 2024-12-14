package main

import (
	"fmt"
	"sync"
)

func main() {
	// 1 is a bufferchannel
	mychan := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, ischannelOpen := <-mychan
		fmt.Println(ischannelOpen)
		fmt.Println(val)
		// fmt.Println(<-mychan)
		// fmt.Println(<-mychan)
		wg.Done()
	}(mychan, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mychan <- 0
		close(mychan)

		wg.Done()
	}(mychan, wg)
	wg.Wait()
}
