package main

import "fmt"

// Write a function that simulates the behavior of the Itoa function in Go. Itoa transforms a number represented as an int in a number represented as a string.

//For this exercise the handling of the signs + or - does have to be taken into account.

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	result := ""
	for n > 0 {

		result = string(rune(n%10+'0')) + result
		n = n / 10
	}
	// Add the negative sign if needed
	if isNegative {
		result = string('-') + result
	}
	return result
}
func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
