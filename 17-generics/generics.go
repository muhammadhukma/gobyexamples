package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int { // as an example of a generic function. SlicesIndex take a slice of any comparable type and an element of that type and return the index of the first occurrence of v in s, or -1 if not present.
	// the comparable constraint means that we can compare values of this type with the == and != operators.
	// for a more thorough explanation of this type signature.
	for i := range s { // note this function exists in the standard library as slice.Index
		if v == s[i] {
			return i
		}
	}
	return -1 // false
}

type List[T any] struct { // as an example of a generic type, list is a singly-linked list with values of any type.
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) { // we can define methods on generic type just like we do on regular types, but we have to keep the type parameters in place.
	if lst.tail == nil { // the type is List[T], not List.
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T { // AllElements returns all the list elements as a slice. In the next example we'll see a more idiomatic way to iterating over all elements of custum types.
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo")) // when invoking generic functionc, we can often rely on type inference. note that we don't have to specify the types for s and E when calling SliceIndex - the compiler infers them automatically.

	_ = SlicesIndex[[]string, string](s, "zoo")
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
