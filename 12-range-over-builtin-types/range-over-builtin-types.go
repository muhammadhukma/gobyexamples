package main

import "fmt"

// range iterates over element in a variety of builtin data structures. Let's see how to use range with some of the data structures we've already learned.

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums { // here we use range to sum the numbers in a slice. arrays work like this too.
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums { // range on arrays and slices provides both the index and value for each entry. above we didn't need the index, so we ignore it with the blank identifier_. sometimes we actually want the indexes though.
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}

	// range can also iterate over just the key of a map.
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k, _ := range kvs {
		fmt.Println("key:", k)
	}
	for _, v := range kvs {
		fmt.Println("value:", v)
	}

	// range on strings iterates over unicode code points.
	// the first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
