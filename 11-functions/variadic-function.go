package main

import "fmt"

func sum(nums ...int) { // here's a function that will take an arbitrary number of ints as arguments.
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums { // within the function, the type of nums is equivalent to []int. we can call len(nums), iterate over it with range, etc.
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2) // variadic functions can be called in the usual way with individual arguments.
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5} // if you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	sum(nums...)

}
