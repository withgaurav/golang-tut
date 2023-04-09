package main

import (
	"fmt"
)

const access string = "ygugwgyweu"

func main() {
	var username string = "gaurav"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("Variable is of type %T \n", isloggedin)

	var val uint8 = 255
	fmt.Println(val)
	fmt.Printf("Variable is of type %T \n", val)

	// default values and some aliases
	var anothervalue int
	fmt.Println(anothervalue)
	fmt.Printf("Variable is of type %T \n", anothervalue)

	// implicit type
	var newval = "gaurav.shukla.in"
	fmt.Println(newval)

	number := 3000.0
	fmt.Println(number)

	fmt.Println(access)
	fmt.Printf("variable is of type : %T \n", access)
}
