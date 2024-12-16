package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	r := 1

	for k := 1; k <= nb; k++ {
		if ((1<<63 - 1) / r) < 0 {
			return 0
		}
		r *= k

	}
	return r
}
