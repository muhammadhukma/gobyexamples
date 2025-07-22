package main

import (
	"fmt"
	"time"
)

// Rate Limiting is an important mechanism for controlling resource utilization and mantaining quiality of service.
// Go elegantly supports rate limiting with go routines, channel, and ticker.

func main() {
	// First we'll look at basic rate limiting. Suppose we want to limit our handling of incoming request. We'll serve these request off a channel of the same name.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond) // this limiter channel will receive a value every 200 millisecond. this regulator in our rate limiting schame.

	// by locking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3) // we may want to allow short bursts of request in our rate limiting schame while preserving the overall rate limit.
	// we can accomplish this by buffering our limiter channel. this burstyLimiter channel will allow bursts of up to 3 events.

	// fill up the channel to represent allowed bursting.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// every 200 millisecond we'll try to add a new value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// now simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability of burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
