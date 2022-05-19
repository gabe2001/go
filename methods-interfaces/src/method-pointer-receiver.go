// (c) 2022 - Hewlett Packard Enterprise Development LP
/*
@author gabe@hpe.com - 2022-05-19


*/
package main

import "fmt"

func main() {

	var i int

	fmt.Println(i)
	dirtySideEffects(&i)
	fmt.Println(i)

}

func dirtySideEffects(i *int) {
	// i += 2 // this would be pointer arithmetic and not actually incrementing int i
	*i += 2
}
