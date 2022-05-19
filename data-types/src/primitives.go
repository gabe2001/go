// (c) 2022 - Hewlett Packard Development L.P.
/*
@author gabe@hpe.com

Demonstrate the use of primitives
  - t is a boolean
  - f is a boolean
  - a is an int, default value 0
  - b is an int too
*/
package main

import "fmt"

func main() {

	t := true
	f := false

	if !(t && f) {
		fmt.Println("this is true")
	} else {
		fmt.Println("this is false")
	}

	var a int

	b := a
	fmt.Println("A is", a, "and B is", b)

	a = 2
	fmt.Println("A is", a, "and B is", b)

	c := &a
	fmt.Println("A is", a, "and C is", *c)

	a = 3
	fmt.Println("A is", a, "and C is", *c)

}
