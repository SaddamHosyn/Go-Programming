package main

import (
	"fmt"
)

//Write a function DigitLen() that takes two integers as arguments and returns the times the first int can be divided by the second until it reaches zero. The second int must be between 2 and 36. If not, the function returns -1. If the first int is negative, reverse the sign and count the digits.

func DigitLen(n, base int) int {
	// Check if base is within the valid range
	if base < 2 || base > 36 {
		return -1
	}

	// If n is negative, reverse its sign
	if n < 0 {
		n = -n
	}

	// Count how many times n can be divided by base until it reaches zero
	count := 0
	for n > 0 {
		n /= base
		count++
	}

	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))  // Output: 3
	fmt.Println(DigitLen(100, 2))   // Output: 7
	fmt.Println(DigitLen(-100, 16)) // Output: 2
	fmt.Println(DigitLen(100, -1))  // Output: -1
}
