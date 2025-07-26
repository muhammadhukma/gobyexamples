package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{1, 4, 2}
	slices.Sort(ints)
	fmt.Println("Ints:  ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("sorted: ", s)
}
