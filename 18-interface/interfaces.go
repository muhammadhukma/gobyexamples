package main

import (
	"fmt"
	"math"
)

type geometry interface { //  here's a basic interface for geometry shapes.
	area() float64
	perim() float64
}

type rect struct {
	width, height float64 // for our example we'll implement this interface on rect and circle types.
}

type circle struct {
	radius float64
}

func (r rect) area() float64 { // to implement an interface in Go, we just need to implement all the methods in the interface
	return r.width * r.height // here we implement geometry on rects.
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// the implementation for circles.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// if a variable has an interface type, then we can call methods that are in the nameed interface
// here a generic measure function taking advantage of this to work on any geometry
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// sometimes it's useful to know the runtime type of an interface value.
// one opetion is using a type assertion as shown here; another is type switch.
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
