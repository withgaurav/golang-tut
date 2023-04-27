package main

import (
	"fmt"
)

func main() {
	var fruitList = []string{"Apple", "Pineapple", "Banana"}

	fruitList = append(fruitList, "Tomato", "Potato", "Onion")
	fmt.Println(fruitList)

}
