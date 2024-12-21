package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for c := '0'; c <= '7'; c++ {
		for d := '1'; d <= '9'; d++ {
			for e := '2'; e <= '9'; e++ {
				if c < d && d < e {
					z01.PrintRune(c)
					z01.PrintRune(d)
					z01.PrintRune(e)

					if c == '7' && d == '8' && e == '9' {
						z01.PrintRune(10)
						continue
					}
					z01.PrintRune(44)
					z01.PrintRune(32)

				}
			}
		}
	}
}
