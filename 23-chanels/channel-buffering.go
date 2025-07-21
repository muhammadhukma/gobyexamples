package main

import "fmt"

func main() {
	message := make(chan string)

	message <- "buffered"
	message <- "channels"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
