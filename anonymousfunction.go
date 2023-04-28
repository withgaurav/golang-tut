package main

import (
	"fmt"
)

func main() {
	x := func(l int, b int) int {
		return l * b
	}

	fmt.Printf("%T \n", x)
	fmt.Println(x(20, 30))
}
