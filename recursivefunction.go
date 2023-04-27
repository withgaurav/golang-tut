// This line indicates that this is the main package, which is required for a Go program.
package main

// This line imports the fmt package, which provides functions for formatting and printing text.
import "fmt"

func factorial(n int) int {

	if n == 1 {

		return 1
	}

	return n * factorial(n-1)
}

// This is the factorial() function, which takes an integer n as input and returns the factorial of that number. It uses recursion to calculate the factorial.

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

// This is the main() function, which is the entry point for the program. It prompts the user to enter a number, scans that number from the standard input, and checks if the number is negative. If the number is negative, it prints an error message. Otherwise, it calls the factorial() function with the entered number as input, stores the result in a variable called result, and prints the result to the console.
