// program to simulate API calls using goroutines managed by wait group
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func simulateAPICall(id int, wg *sync.WaitGroup) {
	// mark the routine done when the function ends
	defer wg.Done()

	duration := time.Duration(rand.Intn(3)+1) * time.Second
	fmt.Printf("API Call %d started, will take %v\n", id, duration)

	time.Sleep(duration)
	fmt.Printf("API Call %d completed\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// create a wait group
	var wg sync.WaitGroup

	numAPICalls := 5
	for i := 1; i <= numAPICalls; i++ {
		wg.Add(1)
		go simulateAPICall(i, &wg)
	}

	// wait for all the goroutines to complete
	wg.Wait()
	fmt.Println("All API calls completed.")
}
