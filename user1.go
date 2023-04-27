package main

import (
	"fmt"
)

func main() {
	var name string
	var is_muggle bool
	fmt.Println("Enter your name & are you a muggle: ")
	fmt.Scanf("%s %t", &name, &is_muggle)
	fmt.Println(name, is_muggle)
}
