package main

import (
	"fmt"
)

//Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.

func CountAlpha(s string) int {
	count := 0 // Initialize a counter for alphabetic characters

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') { // Check if the character is alphabetic
			count++ // Increment the counter
		}
	}

	return count // Return the total count of alphabetic characters
}

func main() {
	fmt.Println(CountAlpha("Hello world")) // Output: 10
	fmt.Println(CountAlpha("H e l l o"))   // Output: 5
	fmt.Println(CountAlpha("H1e2l3l4o"))   // Output: 5
}
