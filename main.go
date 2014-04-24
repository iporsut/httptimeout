package main

import (
    "fmt"
    "net/http"
    "time"
    "errors"
)

func GetTimeout(url string, timeout time.Duration) (*http.Response, error) {
    responseChannel := make (chan *http.Response, 1)
    errorChannel := make (chan error, 1)
    go func() {
        resp, err := http.Get(url)
        if err != nil {
            errorChannel <- err
        } else {
            responseChannel <- resp
        }
    }()

    select {
        case resp := <-responseChannel:
            return resp, nil
        case err := <-errorChannel:
            return nil, err
        case <-time.After(timeout):
            return nil, errors.New("Timeout")
    }
}

func main() {
    resp, err := GetTimeout("https://dl.dropboxusercontent.com/s/t15wh4iw49ekawi/Legend.mp3?dl=1&token_hash=AAF6niX0nqa1GA8d2YaeTpvjtyRl7M1J1TN3goHmAihC3Q", 1*time.Second)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(resp)
    }
}
