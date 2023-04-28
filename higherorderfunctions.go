package main

import (
	"fmt"
)

func calcArea(r float64) float64 {
	return 3.14 * r * r
}
func calcDiameter(r float64) float64 {
	return 2 * r
}
func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}
func main() {
	var query int
	var radius float64
	fmt.Print("enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)
	fmt.Printf("ENter \n 1 - area \n 2- perimeter \n 3 - diamter: ")
	fmt.Scanf("%d", &query)

	if query == 1 {
		fmt.Println("Result: ", calcArea(radius))
	} else if query == 2 {
		fmt.Println("Result: ", calcPerimeter(radius))
	} else if query == 3 {
		fmt.Println("Result: ", calcDiameter(radius))
	} else {
		fmt.Println("Invalid Query")
	}

}
func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("result: ", result)
	fmt.Println("Thankyou!")
}
