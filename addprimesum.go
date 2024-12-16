package main

import (
	"os"

	"github.com/01-edu/z01"
)

//Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or equal to it followed by a newline ('\n').

//If the number of arguments is different from 1, or if the argument is not a positive number, the program displays 0 followed by a newline.

// Checks if a number is prime
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Calculates the sum of all prime numbers less than or equal to n
func sumOfPrimes(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

// Recursively prints digits of a number
func printDigits(n int) {
	if n == 0 {
		return

	}
	printDigits(n / 10)
	z01.PrintRune(rune(n%10 + '0'))
}

// Prints a number by calling printDigits
func printNumber(n int) {
	//if n == 0 {
	//	return
	//}
	printDigits(n)
}

func main() {
	// Check if exactly one argument is provided
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	args := os.Args[1]
	n := 0

	// Validate that the argument is a positive number and not letter or special character.
	for _, r := range args {
		if r < '0' || r > '9' {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
		n = n*10 + int(r-'0') // Convert string to integer
	}
	// check the number. Check if the input is zero or negative.
	if n <= 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	// calling the calculated sum of primes up to n function
	printNumber(sumOfPrimes(n)) // Print the result
	z01.PrintRune('\n')         // Print newline after result
}
