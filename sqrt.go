package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	sqrt := 0
	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			sqrt = i
			break
		}
	}
	return sqrt
}
