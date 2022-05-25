// (c) 2022 - Hewlett Packard Enterprise Development LP
/*
@author gabe@hpe.com - 2022-05-19
*/
package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (p person) correctAge(correction int) {
	p.age = correction
}

func main() {

	l := person{"Gabe", 29}
	fmt.Println(l)
	l.correctAge(99)
	// ha-ha - not saving it!
	fmt.Println(l)

}
