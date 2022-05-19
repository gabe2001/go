// (c) 2022 - Hewlett Packard Enterprise Development LP
/*
@author gabe@hpe.com - 2022-05-19
*/
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {

	c := Celsius(25)
	f := Fahrenheit(75)

	fmt.Printf("%T: %0.f -> %.1f\n", c, c, c.ToFarenheit())
	fmt.Printf("%T: %0.f -> %0.1f\n", f, f, f.ToCelsius())

}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) ToFarenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
