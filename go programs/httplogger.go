package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    url := "https://jsonplaceholder.typicode.com/posts/1"
    
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    fmt.Println("Status:", response.Status)

    fmt.Println("Headers:")
    for key, value := range response.Header {
        fmt.Printf("%s: %s\n", key, value)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("\nJSON Response:")
    fmt.Println(string(body))
}
