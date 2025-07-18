package main

import "fmt"

// in go, an array is a numbered sequence of element of a specific length. In typical go code, slices are much more common; array are useful in some special scenarios.
func main() {
	var a [5]int // here we create an array a that will hold exactly 5 ints. The type of element and length are both part of array's type. by default an array is zero-valued, which for ints mean 0s.
	fmt.Println("emp:", a)

	a[4] = 100 // we can set a value at an index using the array[index] = value syntax, and get value with array[index].
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) // the builtin len returns the length of an array.

	b := [5]int{1, 2, 3, 4, 5} // use this syntax to declare and initialize an array in one line.
	fmt.Println("dcl:", b)

	b = [...]int{2, 3, 4, 5, 6}
	fmt.Println("dcl:", b) // you can also have the compiler count the number of element for you with ...

	b = [...]int{100, 3: 400, 500} // if you specify the index with :, the elements in between will be zeroed.
	fmt.Println("idx", b)

	var twoD [2][3]int // arrays type are one-dimensional, but you can compose types to build multi-dimensional data strucures.
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{ // you can create and initialize multi-dimensional arrays at once too.
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}
