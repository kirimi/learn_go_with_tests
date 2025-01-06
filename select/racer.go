package main

import (
	"errors"
	"net/http"
	"time"
)

var ErrTimeout = errors.New("request timeout")

func Racer(a, b string) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", ErrTimeout
	}
}

func ping(url string) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponceTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
