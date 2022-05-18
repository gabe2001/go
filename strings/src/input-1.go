package main

import "fmt"

func main() {

	var name string

	fmt.Print("Enter your first name: ")
	fmt.Scanf("%s", &name)

	fmt.Printf("Good day %s\n", name)
	
}
