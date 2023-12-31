package main

import (
	"fmt"
	"strconv"
)

/*
1. Write a program to calculate the simple interest 
First-line has the comma-separated values of the Principal, rate and time (in years) respective
*constraints: Round simple interest float value to 2 decimal places
*/

func main() {
	var principal, rate, time float64

	// Input: Principal, Rate, and Time (in years)
	fmt.Print("Enter Principal: ")
	fmt.Scan(&principal)
	fmt.Print("Enter Rate: ")
	fmt.Scan(&rate)
	fmt.Print("Enter Time (in years): ")
	fmt.Scan(&time)

	// Calculate simple interest
	interest := (principal * rate * time) / 100

	// Round interest to two decimal places
	interestRounded := fmt.Sprintf("%.2f", interest)

	// Print the result
	fmt.Println("Simple Interest:", interestRounded)
}


/*
2. Write a program to calculate the area of the circle, First line has a value of the radius of the circle
constraint
1. Use const PI from the package math package
2. Use the Pow function from the math package
3. Round area float value to 2 decimal places
*/

func main() {
	var radius float64

	// Input: Radius of the circle
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&radius)

	// Calculate the area of the circle
	area := math.Pi * math.Pow(radius, 2)

	// Round the area to two decimal places
	areaRounded := math.Round(area*100) / 100

	// Print the result
	fmt.Printf("The area of the circle is %.2f\n", areaRounded)
}
