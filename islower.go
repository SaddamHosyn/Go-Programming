package piscine

func IsLower(s string) bool {
	for _, char := range s {
		if char < 'a' || char > 'z' {
			return false
		}
		if char >= 'A' && char <= 'Z' {
			return false
		}
	}
	return true
}
