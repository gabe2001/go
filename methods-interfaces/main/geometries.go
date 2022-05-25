// (c) 2022 - Hewlett Packard Enterprise Development LP
/*
@author gabe@hpe.com - 2022-05-19

Types, methods, interfaces
 - Data encapsulation
 - Interface describing which methods need to exist for the data objects (structs or aliased primitives)
 - Pointer receiver vs. value receiver methods
   - for side effects we need pointer receivers.
   - we cannot mix painter and value receiver methods in the same "class"
*/
package main

import "fmt"

// interface
type geometry interface {
	volume() float64
	surface() float64
	scale(factor float64)
}

// cubes
type cube struct {
	vertex float64
}

func (c *cube) surface() float64 {
	return 6 * c.vertex * c.vertex
}

func (c *cube) volume() float64 {
	return c.vertex * c.vertex * c.vertex
}

func (c *cube) scale(f float64) {
	c.vertex *= f
}

// bars
type bar struct {
	width, height, length float64
}

func (b *bar) surface() float64 {
	return b.width*b.height*2 + b.height*b.length*2 + b.width*b.length*2
}

func (b *bar) volume() float64 {
	return b.width * b.height * b.length
}

func (b *bar) scale(f float64) {
	b.width *= f
	b.height *= f
	b.length *= f
}

// main
func main() {

	var b = bar{2, 3, 5}

	// %0.f - only integral part of float number
	fmt.Printf("%T - %0.f x %0.f x %0.f (w x h x l) %0.0fm2\n", b, b.width, b.height, b.length, b.surface())
	fmt.Printf("%T - %0.0fm3\n", b, b.volume())

	// double it
	b.scale(2)
	fmt.Printf("%T - %0.f x %0.f x %0.f (w x h x l) %0.0fm2\n", b, b.width, b.height, b.length, b.surface())
	fmt.Printf("%T - %0.0fm3\n", b, b.volume())

	// half it -> undo doubling
	b.scale(0.5)
	fmt.Printf("%T - %0.f x %0.f x %0.f (w x h x l) %0.0fm2\n", b, b.width, b.height, b.length, b.surface())
	fmt.Printf("%T - %0.0fm3\n", b, b.volume())

	b = bar{1.5, 2.5, 4.5}

	// %.2f or %0.2f - include two digits of fractional part
	fmt.Printf("%T - %.2fm2\n", b, b.surface())
	fmt.Printf("%T - %.2fm3\n", b, b.volume())
}
