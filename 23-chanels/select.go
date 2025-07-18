package main

import (
	"fmt"
	"time"
)

/*
go select lets you wait on multiple channel operations.
combining goroutines and channels with select is powerful feture of Go.
*/

func main() {
	/*
		for example we'll select across two channels.
	*/
	c1 := make(chan string)
	c2 := make(chan string)

	// each channel will receive a value after some amount of time, to simulate e.g. blocking rpc operations executing in concurrent goroutines.
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "two"
	}()

	// we'll use select to await both of these values simultaneously, printing each one as it arrives.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("receive", msg1)
		case msg2 := <-c2:
			fmt.Println("receive:", msg2)
		}
	}

}
