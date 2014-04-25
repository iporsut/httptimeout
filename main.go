package main

import (
    "errors"
    "net/http"
    "time"
)

func GetTimeout(url string, timeout time.Duration) (*http.Response, error) {
    responseChannel := make(chan *http.Response, 1)
    errorChannel := make(chan error, 1)
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
