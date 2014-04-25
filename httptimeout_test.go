package httptimeout

import (
    "net/http"
    "testing"
    "time"
)

type MockHttpClient http.Client

func (client *MockHttpClient) Get(url string) (*http.Response, error) {
    time.Sleep(100 * time.Millisecond)
    return &http.Response{}, nil
}

func init() {
    DefaultClient = &MockHttpClient{}
}

func TestGetSuccess(t *testing.T) {
    _, err := GetTimeout("http://google.com", 200*time.Millisecond)
    if err != nil {
        t.Error("Expected Success Actual ", err)
    }
}

func TestGetTimeout(t *testing.T) {
    _, err := GetTimeout("http://timeout.com", 50*time.Millisecond)
    if err == nil || err.Error() != "Timeout" {
        t.Error("Expected Timeout Actual ", err)
    }
}
