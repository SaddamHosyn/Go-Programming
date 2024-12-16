/*package piscine

func NRune(s string, n int) rune {
	if n < 0 || n > len(s) || len(s) == 0 {
		return 0
	}
	return []rune(s)[n-1]
}*/

package piscine

func NRune(s string, n int) rune {
	if n <= 0 || len(s) == 0 {
		return 0
	}

	runes := []rune(s)
	count := 0
	for _, r := range runes {
		count++
		if count == n {
			return r
		}
	}

	return 0
}
