package main

import (
	"fmt"
)

// Write a function that takes two uint representing two strictly positive integers and returns their greatest common divisor. If any of the input numbers is 0, the function should return 0.

//In mathematics, the greatest common divisor (GCD) of two or more integers, which are not all zero, is the largest positive integer that divides each of the integers.

func main() {
	fmt.Println(Gcd(42, 10)) // Output: 2
	fmt.Println(Gcd(42, 12)) // Output: 6
	fmt.Println(Gcd(14, 77)) // Output: 7
	fmt.Println(Gcd(17, 3))  // Output: 1
}

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
