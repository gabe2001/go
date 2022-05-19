package main

import "fmt"

type Person struct {
	name    string
	address Address
}

type Address struct {
	city    string
	country string
}

func main() {

	p := Person{
		name: "Gabe",
		address: Address{
			city:    "Amsterdam",
			country: "The Netherlands",
		},
	}

	fmt.Println(p)

}
