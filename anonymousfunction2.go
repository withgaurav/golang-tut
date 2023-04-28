package main

import (
	"fmt"
)

func main() {
	x := func(l int, b int) int {
		return l * b
	}(20, 30)
	fmt.Printf("%T \n", x)
	fmt.Println(x)
}

// an anonymous function that takes two integer arguments returns their product as an integer. The function is assigned to a variable named
// x
// The type of the variable
// x
//  is printed using the
// Printf
//  function, and the anonymous function is called with arguments 20 and 30 and the result is printed using the
// Println
//  function.
// The output should be
// func(int, int) int
//  and
// 600
