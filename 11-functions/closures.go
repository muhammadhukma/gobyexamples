package main

// go supports anonymous functions, which can form closures.
// anonymous function are useful when you want to define a function inline without having to name it.

import "fmt"

func intSeq() func() int { // this function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq() // we call intSeq. assignning the result (a function) to nextInt. this function value captures its own i value, which will be update each time we call nextInt.

	fmt.Println(nextInt()) // see the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() // to confirm that the state is unique to that particular function, create and test a new one.
	fmt.Println(newInts())
 
