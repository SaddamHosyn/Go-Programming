package piscine

func Any(f func(string) bool, ar []string) bool {
	for _, val := range ar {
		if f(val) {
			return true
		}
	}
	return false
}
