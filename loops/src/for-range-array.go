package main

import "fmt"

func main() {

	strings := []string{
		"The",
		"Dude",
		"abides",
	}

	for index, word := range (strings) {
		fmt.Println("Index:", index, ", word:", word)
	}

}
