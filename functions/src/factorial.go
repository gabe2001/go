// (c) 2022 - Hewlett Packard Enterprise Development LP
/*
@author gabe@hpe.com - 2022-05-19

Recursive function call to calculate factorial(x)
*/
package main

import "fmt"

func main() {

	uints64 := [...]uint64{3, 4, 5, 6, 7, 8, 9, 10, 17, 18, 20, 21}
	floats64 := [...]float64{3, 4, 5, 6, 7, 8, 9, 10, 17, 18, 20, 21}

	for i := 0; i < len(uints64); i++ {
		fmt.Printf("fact(%d) and ffact(%f)\n", uints64[i], floats64[i])
		fmt.Println(fact(uints64[i]))
		fmt.Println(ffact(floats64[i]))
	}

}

func fact(n uint64) uint64 {
	if n < 2 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func ffact(n float64) float64 {
	if n < 2 {
		return 1
	} else {
		return n * ffact(n-1)
	}
}
