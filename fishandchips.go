package main

import (
	"fmt"
)

//Write a function called FishAndChips() that takes an int and returns a string.

//If the number is divisible by 2, print fish.
//If the number is divisible by 3, print chips.
//If the number is divisible by 2 and 3, print fish and chips.
//If the number is negative return error: number is negative.
//If the number is non divisible by 2 or 3 return error: non divisible.

// Note: the order of the condtion is important. If the number is divisible by 2 and 3, it should be placed before n%2 == 0 and n%3 == 0. otherwise the desired output will not be achieved.
func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	if n == 0 {
		return "fish and chips"
	}
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}
	if n%2 == 0 {
		return "fish"
	}
	if n%3 == 0 {
		return "chips"
	}
	return "error: non divisible"
}

func main() {
	fmt.Println(FishAndChips(4))  // fish
	fmt.Println(FishAndChips(9))  // chips
	fmt.Println(FishAndChips(6))  // fish and chips
	fmt.Println(FishAndChips(-5)) // error: number is negative
	fmt.Println(FishAndChips(7))  // error: non divisible
}
