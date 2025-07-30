package main

// go's math/rand/v2 package provides pseudorandom number generation.

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Print(rand.IntN(100), ",") // for example, rand.IntN returns a random int n, 0 <= n < 100.
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
	fmt.Println(rand.Float64())

	// this can be used to generate random floats in other ranges, for example 5.0 <= f' <10.0.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// if you want a known seed, create a new rand.Source and pass it into the new constructor.NewPCG creates a new PCG source that requires a seed of two uint64 numbers.
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
