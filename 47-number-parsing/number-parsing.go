package main

// Parsing numbers from strings is a basic but common task in many programs; here how to do it in go

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64) // the built-in package strconv provides the number parsing.
	fmt.Println(f)                          // with parse float, this 64 tells how many bits of precision to parse.

	i, _ := strconv.ParseInt("123", 0, 64) // for parseInt, the 0 means infer the base from the string, 64 requires that the result fit in 64bits.
	fmt.Println(i)

	// parseInt will recoginze hex-formatted numbers.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// A ParseUint is also available.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi is convenience function for basic base-10 int parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse function return an error on bad input.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
