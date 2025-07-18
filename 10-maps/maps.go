package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int) // to create an empty map, use the builtin make: make(map[key-type]val-type).

	m["k1"] = 7 // set key/value pairs using typical name[key] = val syntax.
	m["k2"] = 13
	_, t := m["k2"]
	fmt.Println("tr", t)

	fmt.Println("map:", m) // printing a map with e.g fmt.Println will show all of its key/value pairs.

	v1 := m["k1"] // get a value for a key with name [key]
	fmt.Println("v1:", v1)

	v3 := m["k3"] // if the value doesn't exits, the zero value of the value type is returned.
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m)) // the builtin len returns the number of key/value pairs when called on a map.

	delete(m, "k2") // the builtin delete removes key/values pairs from a map.
	fmt.Println("map:", m)

	clear(m) // to remove all key/value pairs from a map. use the clear builtin.
	fmt.Println("map:", m)

	_, prs := m["k2"] // the optional second return value when getting a value from a map indicates if the key was present in the map. this can be used to disambiguate between missing keys and keys with zero values like 0 or "". here we didn't need the value itself, so we ignored it with the blank identifier _.
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} // you can also declare and initialize a new map in the same line with those sintax
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2} // the maps package contains a number of useful utility functions for maps.
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
