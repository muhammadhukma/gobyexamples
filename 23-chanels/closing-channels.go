package main

import "fmt"

// closing a channels mengindikasikan bahwa tidak ada lagi nilai yang akan dikirim ke channel.
// ini sangat berguna untuk mengkomunikasikan penyelesaian ke penerima channel.

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true // Send a value to notify that we're done.
				return
			}
		}
	}()

	for j := 1; j <= 9; j++ { // send 3 jobs to the worker over the jobs channels, then closes it.
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done // we await the eorker using the synchronization approach we saw earlier.

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
