package piscine

func IsAlpha(s string) bool {
	for _, char := range s {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') && (char != ' ') {
			return false
		}
	}
	return true
}
