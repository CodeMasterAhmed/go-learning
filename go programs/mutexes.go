// counter program using goroutines and mutexes
package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	// create a mutex
	var mutex sync.Mutex
	var wg sync.WaitGroup

	numGoroutines := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// lock the mutex before changing the counter, so that no other goroutine changes it at same time
			mutex.Lock()
			counter++
			// unlock the mutex after the value is updated
			mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}
