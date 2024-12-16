package piscine

func ToUpper(s string) string {
	chars := []rune(s)

	for i := 0; i < len(chars); i++ {
		if chars[i] >= 'a' && chars[i] <= 'z' {
			chars[i] -= 32
		}
	}

	return string(chars)
}
