package main

import (
	"fmt"
	"sync"
)

// in the previous example we saw how to manage simple counter state using atomic operations.
// for more complex state we can use a mutex to safely access data across multiple goroutines.
// container hold a map of counters; since we want to update it concurrently from multiple goroutines,
// we add a mutex to synchronize access.
/* ______________________________________________________________________________________*/
// Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Lock the mutex before accessing counter; unlock it at the end of the function using a defer statement
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// note that the zero value of a mutex is usable as-is. so no initialization is required here
func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup
	// this function increments a named counter in a loop.
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
		wg.Done()
	}

	// Run several goroutines concurrently;
	// note thtat they all access the same containers, and two of them access the same counter.
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("b", 20000)
	go doIncrement("a", 10000)
	wg.Wait() // wait to finish
	fmt.Println(c.counters)
}
