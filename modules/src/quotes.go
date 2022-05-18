// my main application
package main

import (
	"fmt"
	dude "quotes/lebowski"
	pulp "quotes/pulp-fiction"
)

func main() {
	fmt.Println(dude.GetQuote())
	fmt.Println(pulp.GetQuote())
}
