package main

import (
	"fmt"
)

//Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.

func CheckNumber(arg string) bool {
	for _, char := range arg {
		if char >= '0' && char <= '9' { // Check if the character is between '0' and '9'
			return true // Return true if a digit is found
		}
	}
	return false // Return false if no digits are found
}

func main() {
	fmt.Println(CheckNumber("Hello"))  // Output: false
	fmt.Println(CheckNumber("Hello1")) // Output: true
}
