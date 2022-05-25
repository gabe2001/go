// (c) 2022 - Hewlett Packard Enterprise Development LP
/*
@author gabe@hpe.com

Closures have access to variables outside if their lexical scope

Could this be used to implement a stream-like function? Operate on each item in a list (array, slice)
*/
package main

import "fmt"

func main() {

	var n int

	incr := func() int {
		n += 1
		return n
	}

	fmt.Println(incr()) // 1
	fmt.Println(incr()) // 2
	fmt.Println(incr()) // 3

}
