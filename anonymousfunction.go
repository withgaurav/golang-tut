// An anonymous function is just like functions but it is declared without having name identifier to refer it.
// They can also return outputs and take inputs.
// Let's take an example
package main

import (
	"fmt"
)

// Inside our main function we have declared an anonymous function so it is taking two integers as an integer and returning product of those integers as an integer.
func main() {
	// SO note it does not any name specfied but it is returning an integer after the func keyword. And well say this keyword in the variable "x".
	x := func(l int, b int) int {
		return l * b
	}

	fmt.Printf("%T \n", x)
	// let's print the datatype of this variable.
	fmt.Println(x(20, 30))
	// Well! Lets call x with argument
}

// FInally when you'll run this program you will get the product of 20 and 30 which is 600
