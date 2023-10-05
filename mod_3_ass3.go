
package main

import (
	"fmt"
)
/*
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
*/

/* 
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

