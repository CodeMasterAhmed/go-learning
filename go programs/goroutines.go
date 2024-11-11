// program to check the status of a website using go routines
package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s is down!\n", url)
		return
	}
	fmt.Printf("%s is up! Status Code: %d\n", url, resp.StatusCode)
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.invalidurltest.com",
	}

	for _, url := range urls {
		// the function is called using go routine to start a separate thread
		go checkStatus(url)
	}

	time.Sleep(2 * time.Second)
}
