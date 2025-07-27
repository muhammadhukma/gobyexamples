package main

import (
	"cmp"
	"fmt"
	"slices"
)

// sometimes we'sll wnat to sort a collection by something other than its natural order.
// For example, suppose we wanted to sort strings by their length instead of alphabetically.
// Here's an example of cusxtom sort in Go.

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// we emplement a comparison function for string length. cmp.Compare is helpful for this.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}
