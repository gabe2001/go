// (c) 2022 - Hewlett Packard Enterprise Development LP
/*
@author gabe@hpe.com - 2022-05-19
*/
package main

import "fmt"

func main() {

	var ints = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(ints)
	fmt.Println(fReverse(ints))
	fmt.Println(ints)

}

func fReverse(list []int) []int {
	var rList []int
	for i := len(list) - 1; i > -1; i-- {
		rList = append(rList, list[i])
	}
	return rList
}
