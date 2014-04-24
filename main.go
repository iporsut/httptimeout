package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    responseChannel := make (chan *http.Response, 1)
    errorChannel := make (chan error, 1)
    go func() {
        resp, err := http.Get("google.com/")
        if err != nil {
            errorChannel <- err
        } else {
            responseChannel <- resp
        }
    }()

    select {
        case resp := <-responseChannel:
            fmt.Println(resp)
        case err := <-errorChannel:
            fmt.Println(err)
        case <-time.After(1*time.Second):
            fmt.Println("Timeout")
    }
}
