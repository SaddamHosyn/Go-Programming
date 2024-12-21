package piscine

func ToLower(s string) string {
	chars := []rune(s)

	for i := 0; i < len(chars); i++ {
		if chars[i] >= 'A' && chars[i] <= 'Z' {
			chars[i] += 32
		}
	}

	return string(chars)
}
