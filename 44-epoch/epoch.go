package main

import (
	"fmt"
	"time"
)

/*
A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch, Here's ho to do it in Go.
*/

func main() {
	now := time.Now() // use time.Now() with unix. uniMlli or UniNano to get elapsed time since Unix epoch in seconds, milliseconds, or nanoseconds, respectively.
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	//you can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
