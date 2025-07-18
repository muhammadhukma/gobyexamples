package main

import "fmt"

// go supports methods defined on struct types.

type rect struct {
	width, height int
}

func (r *rect) area() int { // this area method has a reciever type of *rect.
	return r.width * r.height
}

func (r rect) perim() int { // methods can be defined for either pointer or value receiver types. here's an example of a value receiver.
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}

	fmt.Println("area: ", r.area()) // here we call the 2 methods defined for our struct
	fmt.Println("perim: ", r.perim())

	rp := &r                         // go automatically handles conversion between values and pointers for method calls.
	fmt.Println("area: ", rp.area()) // you may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	fmt.Println("perim: ", rp.perim())
}
