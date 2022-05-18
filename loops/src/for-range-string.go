package main

import "fmt"

func main() {

	str := "Hello \U0001F972 !"

	fmt.Println(str)

	for index, char := range str {
		// notice the skipping of the UTF-8 non-printable stuff from the string's smiley face icon
		fmt.Printf("Index: %2d - char %c - code %3d\n", index, char, char)
	}

}
