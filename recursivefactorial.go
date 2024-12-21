package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 2 {
		return 2
	}
	if nb == 1 || nb == 0 {
		return 1
	}
	return RecursiveFactorial(nb-1) * nb
}
