// program to simulate communication between multiple goroutines and a go channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// create a channel to recieve the logs
	logChannel := make(chan string)

	go func() {
		// print the logs recieved by goroutines in go channel
		for log := range logChannel {
			fmt.Println("Log:", log)
		}
	}()

	for i := 1; i <= 3; i++ {
		go func(taskID int) {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// send the log to go channel
			logChannel <- fmt.Sprintf("Task %d completed", taskID)
		}(i)
	}

	time.Sleep(2 * time.Second)

	close(logChannel)
}
