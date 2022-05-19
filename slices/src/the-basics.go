/* (c) 2022 - Hewlett Packard Enterprise Development
   @author: gabe@hpe.com
*/
package main

import "fmt"

func main() {

	var x = [...]int{10, 20, 30} // this as an array
	var y = []int{10, 20, 30}    // this is a slice

	var z = []int{1, 5: 4, 6, 10: 100, 15} // a slice with 12 entries, mostly zeros

	var ticTacToe [][]int // a slice of slices

	fmt.Println(x, y, z, ticTacToe)

	fmt.Printf("%d - %s\n", len(x), x)
	fmt.Printf("%d - %s\n", len(y), y)
	fmt.Printf("%d - %s\n", len(z), z)

	fmt.Printf("%d - %s\n", len(ticTacToe), ticTacToe)

	//
	ints := make([]int, 3, 6)
	fmt.Printf("%d - %s\n", len(ints), ints)

	for i, x := range (ints) {
		fmt.Printf("%d - %x\n", i, x)
	}

}
