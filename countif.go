package piscine

func CountIf(f func(string) bool, tab []string) int {
	nb := 0
	for _, val := range tab {
		if f(val) {
			nb++
		}
	}
	return nb
}
