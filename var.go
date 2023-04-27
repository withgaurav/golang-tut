// variable scope
// In golang scope is defined as block. A block is a possibly empty sequence of declaration and statements within matching curly bracke
// In golang there are two types of block inner block and outer block
// Inner block can access the variables that are declared in the outer block
// Outer block cannot access the variables that are declared in the inner block.
package main

import "fmt"

func main() {
	city := "London"
	{
		country := "UK"
		fmt.Println(city)
		fmt.Println(country)
	}
	fmt.Println(city)
}
