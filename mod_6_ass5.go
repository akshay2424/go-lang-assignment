package main

import {
	"errors"
	"fmt"
	}

/* 
1. In the given code, the accessSlice function accepts a slice and index. 
If the value is present in the slice at that index, the program should print the following.

"item <index>, value <value at that index in slice>"

But if the index does not hold any value,
it will lead to an index out of range panic in our program.

Complete the given code to recover from panic inside the accessSlice function, when index out of range panic occurs.
Also, Print the following after handling panic

"internal error: <recovered panic value>"


Example Test Case 1 :

Input: 3
Output:
item 3, value 6
Testing Panic and Recover
*/

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 10

	accessSlice(slice, index)
	fmt.Println("Testing Panic and Recover")
}

func accessSlice(slice []int, index int) {
	if index >= 0 && index < len(slice) {
		value := slice[index]
		fmt.Printf("item %d, value %d\n", index, value)
	} else {
		fmt.Printf("internal error: index out of range\n")
	}
}


/*
2. In the given code, the accessSlice function accepts a slice and index.
If the value is present in the slice at that index, the program should print the following.

"item <index>, value <value present at the index>"

But if the index does not hold any value,
it will lead to an index out of range panic in our program.


So in order to safeguard our program from panicking, add a condition to handle the condition if 

index > lengthOfSlice - 1

and return an error from the accessSlice function if the above condition is met.
The error message should be the following :

"length of the slice should be more than index"

Complete the given program to return an error from the accessSlice function and handle the returned error inside the main function to print the message. 


Example Test Case 1 :

Input: 3
Output:
item 3, value 6
*/
func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 10

	value, err := accessSlice(slice, index)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("item %d, value %d\n", index, value)
	}
}

func accessSlice(slice []int, index int) (int, error) {
	if index < 0 || index >= len(slice) {
		return 0, errors.New("length of the slice should be more than index")
	}

	return slice[index], nil
}

