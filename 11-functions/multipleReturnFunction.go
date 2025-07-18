package main

import "fmt"

// go has built-in support for multiple return values. this feature is used often in idiomatic go, for example to return both result and error values from a function.

func vals() (int, int) { // the (int, int) in this function signature shows that the function return 2 ints.
	return 3, 7
}

func main() {
	a, b := vals() // here we use the 2 defferent return values from the call with multiple assignment.
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // if you only want a subset of the returned, use the blank identifier _.
	fmt.Println(c)
}
