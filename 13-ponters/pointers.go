package main

import "fmt"

// go supports pointers. allowing you to pass references to value and records within your program.

// we'll show how pointes work in contrast to values with 2 functions: zeroval and zeroptr.
// zeroval has an int parameter, so arguments will be passed to it by value.
// zeroptr will get a copy of ival distinct from the one in the calling function.

func zeroval(ival int) {
	ival = 0
}

// zeroptr in contras has an *int parameter, meaning that it takes an int pointer.
// the *iptr code in the function body then deferences the pointer from its memory address to the current value at that address.
// assigning a value to a dereferenced pointer changes the value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// the &i syntax gives the memory address of i, i.e a ponter to i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// pointer can be printed too
	fmt.Println("pointer:", &i)
}
