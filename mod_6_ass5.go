package main

import {
	"errors"
	"fmt"
	}

/* 1 que
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
*/

//2 nd 
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

