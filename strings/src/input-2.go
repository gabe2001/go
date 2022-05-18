package main

import "fmt"

// note Scanf is reading the input string word by word (whitespace separated)
func main() {

	var name, last string

	fmt.Print("Enter your first name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter your  last name: ")
	fmt.Scanf("%s", &last)

	fullName := name + " " + last
	fmt.Printf("Good day %s\n", fullName)

}
