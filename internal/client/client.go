package client

import (
	"net/http"
	"time"
)

func DoRequest(url string) (int, time.Duration, error) {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	if err != nil {
		return 0, duration, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, duration, nil
}
