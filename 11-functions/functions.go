package main

import "fmt"

func plus(a int, b int) int { // here's a function that takes two ints and returns their sum as an int.
	return a + b // go requires explicit return, i.e it won;t automatically return the value of the last expression.
}

// when you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declare the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
