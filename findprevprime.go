package main

import "fmt"

//Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

//If there are no primes inferior to the int passed as parameter the function should return 0.

func isPrime(n int) bool {
	if n < 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}

	}
	return 0
}

func main() {
	fmt.Println(FindPrevPrime(5)) // Output: 5
	fmt.Println(FindPrevPrime(4)) // Output: 3
}
