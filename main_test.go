package main

import (
    "testing"
    "time"
)

func TestGetSuccess(t *testing.T) {
    _, err := GetTimeout("http://google.com", 5*time.Second)
    if err != nil {
        t.Error("Expected Success Actual ", err)
    }
}

func TestGetError(t *testing.T) {
    _, err := GetTimeout("http://localhost.com", 1*time.Second)
    if err == nil && err.Error() == "Timeout" {
        t.Error("Expected Error Actual ", err)
    }
}

func TestGetTimeout(t *testing.T) {
    _, err := GetTimeout("https://dl.dropboxusercontent.com/s/t15wh4iw49ekawi/Legend.mp3?dl=1&token_hash=AAF6niX0nqa1GA8d2YaeTpvjtyRl7M1J1TN3goHmAihC3Q", 1*time.Second)
    if err == nil || err.Error() != "Timeout" {
        t.Error("Expected Timeout Actual ", err)
    }
}
