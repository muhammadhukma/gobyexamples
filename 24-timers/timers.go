package main

// we often want to execute go code at some point in the future, or repeatedly at some interval.
// go built-in timer and ticker feature make both of these tasks easy.
// we will look first at timers and then at ticker.

import (
	"fmt"
	"time"
)

func main() {
	// timers represent a single event in the future. you tell the timer how long you wnat to wait, and it provides a channel that will be notified at that time.
	// this timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// the <-timer1.C bloack on the timer's channel C until it send a value indicating that the timer fired.
	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
