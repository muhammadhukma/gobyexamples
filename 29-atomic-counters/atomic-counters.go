package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// the primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools.
// There are a few other options for managing state though.
// here we'll look at using the sync/atomic package for atomic counters accessed by multiple goroutines'

func main() {
	var ops atomic.Uint64 // we'll use an atomic integer type to represent our (always-positive) counter.

	var wg sync.WaitGroup // a WaitGroup will help us wait for all goroutines to finish their work.

	for range 50 { // We'll start 50 goroutines that each increment the counter exactly 1000 times.
		wg.Add(1)

		go func() {
			for range 1000 {
				ops.Add(1) // to atomic increment the counter we use add.
			}
			wg.Done()
		}()
	}
	wg.Wait()                       //wait until all the goroutines are donel.
	fmt.Println("ops:", ops.Load()) // here no goroutines are writing to 'ops', but using Load its safe to atomatically read a value even while other goroutines are (atomatically) updating it.
}
