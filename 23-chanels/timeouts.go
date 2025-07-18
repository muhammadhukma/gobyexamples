package main

import (
	"fmt"
	"time"
)

/*
* Timeouts are importan for program that connect to external resources or that otherwise need to bound execution time
* Implementing timeouts in Go is easy and elegant thanks to channels and select.
 */

func main() {
	// For example, suppose we're executing an external call that returns its result on a channel c1 after 2s.
	// Note that the channel is buffered, so the send in goroutine is nonblocking.
	// This is a common pattern to prevent goroutine leaks in case the channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
