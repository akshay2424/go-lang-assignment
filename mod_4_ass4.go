package main

import (
	"fmt"
)

/*

1. The given program accepts an integer value between 1 to 4 from the user.
There is an option associated with each value, which is basically objects of different data types with some associated value.

Write a method named "AcceptAnything" which takes any type of data as input.

Based on the option chosen by the user, we will send different types of objects to this function and it should print the following based on its respective data type of value sent to the function "AcceptAnything".

integer :
"This is a value of type Integer, <value>"

string :
"This is a value of type String, <value>"

boolean :
"This is a value of type Boolean, <value>"

Custom data type Hello :
"This is a value of type Hello, <value>"
*/
// Function to handle and print different data types
func AcceptAnything(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("This is a value of type Integer: %d\n", v)
	case string:
		fmt.Printf("This is a value of type String: %s\n", v)
	case bool:
		fmt.Printf("This is a value of type Boolean: %t\n", v)
	default:
		fmt.Println("Unsupported data type")
	}
}

func main() {
	// Prompt the user for input
	fmt.Println("Choose an option:")
	fmt.Println("1. Integer")
	fmt.Println("2. String")
	fmt.Println("3. Boolean")
	fmt.Println("4. Custom Hello type")
	fmt.Print("Enter your choice (1-4): ")

	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil || choice < 1 || choice > 4 {
		fmt.Println("Invalid choice. Please choose between 1 and 4.")
		return
	}

	switch choice {
	case 1:
		var num int
		fmt.Print("Enter an integer value: ")
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Invalid input for an integer.")
			return
		}
		AcceptAnything(num)
	case 2:
		var str string
		fmt.Print("Enter a string value: ")
		_, err := fmt.Scan(&str)
		if err != nil {
			fmt.Println("Invalid input for a string.")
			return
		}
		AcceptAnything(str)
	case 3:
		var flag bool
		fmt.Print("Enter a boolean value (true or false): ")
		_, err := fmt.Scan(&flag)
		if err != nil {
			fmt.Println("Invalid input for a boolean.")
			return
		}
		AcceptAnything(flag)
	case 4:
		type Hello struct {
			Message string
		}
		var helloObj Hello
		fmt.Print("Enter a message for the Hello type: ")
		_, err := fmt.Scan(&helloObj.Message)
		if err != nil {
			fmt.Println("Invalid input for the Hello type.")
			return
		}
		AcceptAnything(helloObj)
	}
}


/*
2. The given program accepts 2 values from the user, length and breadth of a rectangle respectively.

Complete the program by writing methods "Area" and "Perimeter" on Rectangle type to calculate the area and perimeter of a rectangle.

Hint: 
Method Signatures for Area and Perimeter
Area() int
Perimeter() int

Formulae:
Area = length * width
Perimeter = 2 * (length + width)

Example Test Case 1:
Input: 10 20
Output:
Area of Rectangle: 200
Perimeter of Rectangle: 60

3. The given program takes an integer value as input from the user i.e 1 or 2.
Option 1 represents Rectangle and Option 2 represents Square.

Given the metrics of the shapes are hard-coded, complete the given program to achieve the following:

1. Create an interface Quadrilateral which has the following method signatures
* Area() int
* Perimeter() int

2. Implement the Quadrilateral interface for the given shapes i.e Rectangle and Square
*/

// Define a Rectangle struct
type Rectangle struct {
	Length  int
	Breadth int
}

// Method to calculate and return the area of the rectangle
func (r Rectangle) Area() int {
	return r.Length * r.Breadth
}

// Method to calculate and return the perimeter of the rectangle
func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Breadth)
}

func main() {
	// Prompt the user for input
	fmt.Print("Enter the length of the rectangle: ")
	var length int
	_, err := fmt.Scan(&length)
	if err != nil {
		fmt.Println("Invalid input for length.")
		return
	}

	fmt.Print("Enter the breadth of the rectangle: ")
	var breadth int
	_, err = fmt.Scan(&breadth)
	if err != nil {
		fmt.Println("Invalid input for breadth.")
		return
	}

	// Create a rectangle instance
	rect := Rectangle{
		Length:  length,
		Breadth: breadth,
	}

	// Calculate and print the area and perimeter
	area := rect.Area()
	perimeter := rect.Perimeter()

	fmt.Printf("Area of Rectangle: %d\n", area)
	fmt.Printf("Perimeter of Rectangle: %d\n", perimeter)
}


/*
3. Define a "Print" function which accepts any shape that implements Quadrilateral interface and Prints Area and Perimeter of shape in the following manner:

"Area :  <value>"
"Perimeter :  <value>"


HINT: Step 2 means, to define methods in Quadrilateral interface for given shapes


Formulae:

Area of Rectangle: length * width
Perimeter of Rectangle: 2*(length + breadth)


Area of Square: side * side
Perimeter of Square: 4 * side
*/
// Define an interface Quadrilateral
type Quadrilateral interface {
	Area() int
	Perimeter() int
}

// Define a Rectangle struct
type Rectangle struct {
	Length  int
	Breadth int
}

// Implement the Quadrilateral interface for Rectangle
func (r Rectangle) Area() int {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Breadth)
}

// Define a Square struct
type Square struct {
	Side int
}

// Implement the Quadrilateral interface for Square
func (s Square) Area() int {
	return s.Side * s.Side
}

func (s Square) Perimeter() int {
	return 4 * s.Side
}

// Define a Print function for any shape implementing Quadrilateral interface
func Print(shape Quadrilateral) {
	fmt.Println("Area:", shape.Area())
	fmt.Println("Perimeter:", shape.Perimeter())
}

func main() {
	fmt.Println("Choose a shape:")
	fmt.Println("1. Rectangle")
	fmt.Println("2. Square")
	fmt.Print("Enter your choice (1 or 2): ")

	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil || (choice != 1 && choice != 2) {
		fmt.Println("Invalid choice. Please choose 1 for Rectangle or 2 for Square.")
		return
	}

	if choice == 1 {
		// Rectangle
		rect := Rectangle{Length: 10, Breadth: 20}
		fmt.Println("Rectangle Metrics:")
		Print(rect)
	} else if choice == 2 {
		// Square
		square := Square{Side: 15}
		fmt.Println("Square Metrics:")
		Print(square)
	}
}