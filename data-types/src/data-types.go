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

}
