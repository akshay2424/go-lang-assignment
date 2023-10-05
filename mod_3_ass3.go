
package main

import (
	"fmt"
)
/*
1. Find the Index of the element
Write a program to store the day(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday) against the respective index of the day(1, 2, 3, 4, 5, 6, 7) in a map.
You can hard-code the map in your program.
Take an integer input from the user and print the day stored against that index and if nothing is to be found for that index,
Print "Not a day"
Hint: Following code can be used to take an integer input from the user in the GO Programming Language
var index int
fmt.Scanln(&index)
Example Test Case:
Input: 5
Output: Friday
Input: 11
Output: Not a day
*/

func main() {
	// Create a map to store the days of the week
	daysMap := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	// Prompt the user for input
	var index int
	fmt.Print("Enter an index (1-7): ")
	_, err := fmt.Scanln(&index)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// Check if the index is valid and retrieve the corresponding day
	day, found := daysMap[index]
	if found {
		fmt.Println(day)
	} else {
		fmt.Println("Not a day")
	}
}


/* 
2. Word Count
Write a Program to fulfil the following conditions.
Input: A string containing words separated by space
Output: A slice containing words with the highest frequency in the input string.
Note: The order of the words in the resultant slice should be the same as the order they appear in the input string.
Condition: Words are case-sensitive. Joe is not the same as joe.
Given Code Template: A code template is provided that takes a string as an input and turns it into a slice of strings.
Example Test Case 1:
Input: My name is Joe and My Father's name is also Joe
Output: [My name is Joe]
Here, the words "My", "name", "is", "Joe" appeared 2 times each, which is also the highest frequency of any word.
Hence the output contains all 4 words.
Example Test Case 2:
Input: Europe is far but the US is farther
Output: [is]
Here, the word "is" appeared twice which is also the highest frequency of any word.
*/

func main() {
	// Prompt the user for input
	fmt.Print("Enter a string of words separated by space: ")
	var input string
	fmt.Scanln(&input)

	// Split the input string into words
	words := strings.Fields(input)

	// Create a map to store word frequencies
	wordFreq := make(map[string]int)

	// Find the word with the highest frequency
	maxFrequency := 0

	for _, word := range words {
		wordFreq[word]++
		if wordFreq[word] > maxFrequency {
			maxFrequency = wordFreq[word]
		}
	}

	// Create a slice to store words with the highest frequency
	var result []string
	for _, word := range words {
		if wordFreq[word] == maxFrequency {
			result = append(result, word)
		}
	}

	// Print the result
	fmt.Println(result)
}

/*
3. Return the slices
Complete the program to return 3 slices of a given array, based on the following conditions.
Given Array : ["qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"]
Hint: Hard-code the given array into your program.
Input: Two space-separated integers representing index1 and index2.
Output: Output will contain 3 slices
1. slice containing all the elements from the start of the array to Index1
2. slice containing all the elements from index1 to index2
3. slice containing all the elements from index2 to the end of the array
Conditions to Handle:
If Either of the input indexes is out of range of the given array, print "Incorrect Indexes" and exit the program
Example Test Case 1:
Input: 2 4
Output:
[qwe wer ert]
[ert rty tyu]
[tyu yui uio iop]
Example Test Case 2:
Input: 2 8
Output: Incorrect Indexes
*/
func main() {
	// Hard-coded array
	givenArray := []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

	// Prompt the user for input
	fmt.Print("Enter two space-separated integers (index1 and index2): ")
	var index1, index2 int
	_, err := fmt.Scanf("%d %d", &index1, &index2)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// Check if the indexes are within the range of the array
	if index1 < 0 || index1 >= len(givenArray) || index2 < 0 || index2 >= len(givenArray) {
		fmt.Println("Incorrect Indexes")
		return
	}

	// Create three slices based on the input indexes
	slice1 := givenArray[:index1+1]
	slice2 := givenArray[index1 : index2+1]
	slice3 := givenArray[index2:]

	// Print the three slices
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

