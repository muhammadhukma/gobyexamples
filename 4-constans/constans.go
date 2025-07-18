package main

import (
	"fmt"
	"math"
)

const s string = "constans" // const declare a constan value

func main() {
	fmt.Println(s)

	const n = 500000000 // a constan statement can appear anywhere a var statement can

	const d = 3e20 / n // constan expression perform arithmetic with arbitary precision.
	fmt.Println(d)

	fmt.Println(int64(d)) // a numeric constant has no type until it's given one, such as by an explicit conversion.

	fmt.Println(math.Sin(n)) // a number can be given a type by using it in a context that require one. such as a variable assignment or function call. For example, here math.Sin expect a float64.
}
