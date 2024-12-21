package piscine

func Map(f func(int) bool, ar []int) []bool {
	t := make([]bool, len(ar))

	for m, n := range ar {
		t[m] = f(n)
	}
	return t
}
