package main

import "fmt"

func factorial(n int) int {

	if n == 1 {

		return 1
	}

	return n * factorial(n-1)
}

func main() {

	var n int

	fmt.Print("Enter the Number: ")

	fmt.Scan(&n)

	if n < 0 {

		fmt.Print("\nFactorial of a negative number is not possible")

	}

	result := factorial(n)

	fmt.Printf("\nFactorial is %d", result)

}
