// (c) 2022 - Hewlett Packard Enterprise Development LP
// @author gabe@hpe.com - 2022-05-25
package main

import (
	"fmt"
	float "main/float-math"
	int "main/int-math"
)

func main() {

	fmt.Println(int.Add(2, 1))
	fmt.Println(int.Subtract(2, 1))

	fmt.Printf("%.1f\n", float.Add(2.0, 1.0))
	fmt.Printf("%.1f\n", float.Subtract(2.0, 1.0))

}
