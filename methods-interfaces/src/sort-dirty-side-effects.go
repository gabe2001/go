// (c) 2022 - Hewlett Packard Enterprise Development LP
/*
@author gabe@hpe.com - 2022-05-19
*/
package main

import "fmt"

func main() {

	var ints = []int{1, 2, 3, 4, 5}
	fmt.Println(ints)
	reverse(ints)
	fmt.Println(ints)
}

func reverse(list []int) {
	var max = len(list) - 1
	for i := 0; i < len(list)/2; i++ {
		var x = list[i]
		list[i] = list[max-i]
		list[max-i] = x
	}
}
