package main

import "fmt"

// go's structs are typed collection of fields. they're useful for grouping data together to form records.

type person struct { // this person struct type has name and age fields.
	name string
	age  int
}

func newPerson(name string) *person { // go is a garbage collected language; you can safely return a pointer to local variable.
	p := person{name: name} // it will only be cleaned up by the garbage collector when there are no active reference to it.
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"bob", 20})              // this syntax creates a new struct.
	fmt.Println(person{name: "alice", age: 43}) // you can name the fields when initializing a struct.
	fmt.Println(person{name: "freed"})          // ommited field will be zero-valued.
	fmt.Println(&person{name: "Ann", age: 40})  // an & prefix yields a pointer to the struct.
	fmt.Println(newPerson("John"))              // It's idiomatic to encapsulate new struct creation in constructor functions.

	s := person{name: "sean", age: 50} // access struct field with a dot.
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // you can also use dot with struct pointers. the pointer are automically dereferenced.

	sp.age = 51 // structs are mutable.

	dog := struct { // if  a struct type is only used for a single value, we don't have to give it a name.
		name   string // the value can have an anonymous struct type. this technique is commonly used for table-driven tests.
		isGood bool
	}{
		"rex",
		true,
	}
	fmt.Println(dog)
}
