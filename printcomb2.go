package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	tilt := true
	for c := 0; c <= 9; c++ {
		for d := 0; d <= 9; d++ {
			for e := 0; e <= 9; e++ {
				for f := 0; f <= 9; f++ {
					if (c < e) || (c == e && d < f) {
						if !tilt {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
						z01.PrintRune(rune('0' + c))
						z01.PrintRune(rune('0' + d))
						z01.PrintRune(' ')
						z01.PrintRune(rune('0' + e))
						z01.PrintRune(rune('0' + f))
						tilt = false

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
