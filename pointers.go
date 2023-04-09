// Pointer is just the reference to the memory address since you're directly passing on to that address it makes a 100% guarantee that actual value is passed on.
package main

import (
	"fmt"
)

func main() {
	// var ptr *int
	// fmt.Println("The actual number of pointer is: ", ptr)

	myNumber := 35
	// For reference to the memory address we use "&" this
	var ptr = &myNumber
	fmt.Println("For printing the actual value of pointer here is: ", ptr)
	fmt.Println("For printing the actual value of pointer here is: ", *ptr)
	*ptr = *ptr + 2
	fmt.Println("The new value is: ", myNumber)

}
