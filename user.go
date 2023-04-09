package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hey there, ", name)
}
