// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	fmt.Println("Race condition - ")
// 	wg := &sync.WaitGroup{}

// 	var score = []int{0}
// 	wg.Add(3)
// 	go func(wg *sync.WaitGroup) {
// 		fmt.Println("1st race condition")
// 		score = append(score, 1)
// 		defer wg.Done()
// 	}(wg)

// 	go func(wg *sync.WaitGroup) {
// 		defer wg.Done()
// 		fmt.Println("2nd race condition")
// 		score = append(score, 2)
// 	}(wg)

// 	go func(wg *sync.WaitGroup) {
// 		defer wg.Done()
// 		fmt.Println("3rd race condition")
// 		score = append(score, 3)
// 	}(wg)

// 	wg.Wait()
// 	fmt.Println(score)
// }

package main

import (
	"fmt"
	"sync"
)

// This program has a race condition because multiple goroutines
// update the shared variable 'counter' concurrently without synchronization.

func main() {
	var counter int
	var wg sync.WaitGroup

	// Number of goroutines
	const numGoroutines = 10

	wg.Add(numGoroutines)

	// Start multiple goroutines
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++ // Race condition: multiple goroutines modifying 'counter' concurrently
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter)
}
	