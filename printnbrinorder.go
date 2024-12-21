package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	var digits [10]int
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	for n > 0 {
		digits[n%10]++
		n /= 10
	}

	for i := 0; i <= 9; i++ {
		for digits[i] > 0 {
			z01.PrintRune(rune('0' + i))
			digits[i]--
		}
	}
}
