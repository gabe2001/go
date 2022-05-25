// (c) 2022 - Hewlett Packard Enterprise Development LP
/*
@author gabe@hpe.com - 2022-05-25
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Filename not provided")
	}

	log.Print(args)

	file, err := os.Open(args[0])

	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file", file, err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	m := make(map[string]int)

	for scanner.Scan() {
		s := scanner.Text()
		v, existed := m[s]
		if existed {
			m[s] = v + 1
		} else {
			m[s] = 1
		}
	}

	for k, v := range m {
		_, err := fmt.Println(k, ":", v)
		if err != nil {
			return
		}
	}
}
