package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
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
	a = animal(p)

	fmt.Println(a)

	var b animal

	b.name = p.name
	b.age = p.age

	fmt.Println(b)

}
