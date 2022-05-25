package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
		//nickname string // not cast-able due to extra property
	}

	type animal struct {
		name string
		age  int
	}

	p := person{
		name: "Gabe",
		age:  29,
	}

	var a animal
	//a = p // compiler no like!
	a = animal(p) // cast person to animal

	fmt.Println(a)

	var b animal

	b.name = p.name
	b.age = p.age

	fmt.Println(b)

}
