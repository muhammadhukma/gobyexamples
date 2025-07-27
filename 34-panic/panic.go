package main

import "os"

// A panic typically means somehting went unexpectedly wrong.
// Most we use it to fail fast on errors that shouldn't occur during normal operation, or that we arent prepared to handle gracefully.

func main() {
	// We'll use panic throughout this site to check for unexpected errors. this is the only program on the site designed to panic.
	panic("a problem")

	// a common use of panic is to abort if a function return an error value that we don't know how to (or want to) handle. Here's an example of panickingif we get an unexpected error when creating a new file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
