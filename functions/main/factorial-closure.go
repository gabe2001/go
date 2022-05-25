// (c) 2022 - Hewlett Packard Enterprise Development LP
/*
@author gabe@hpe.com

Closures have access to variables outside if their lexical scope

Could this be used to implement a stream-like function? Operate on each item in a list (array, slice)
*/
package main

import "fmt"

// mul multiply its current value with parameter
func mul() func(i int) int {
	x := 1
	return func(i int) int {
		x *= i
		return x
	}
}

func fact(i int) int {
	if i < 2 {
		return 1
	} else {
		return i * fact(i-1)
	}
}

func main() {

	// assign function mul, which contains a closure, to f
	f := mul()
	// call f(i) (n-1) times to calculate factorial(n-1) - when started from 1 or 2
	for i := 2; i < 10; i++ {
		f(i)
	}
	fmt.Println(f(1))

	// and the recursive result will be the same
	fmt.Println(fact(9))

}
