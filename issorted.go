package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	r1 := 1
	r2 := 1
	r3 := 1
	for m, n := range a {
		if m != len(a)-1 {
			if f(n, a[m+1]) < 0 {
				r1++
			}
			if f(n, a[m+1]) > 0 {
				r2++
			}
			if f(n, a[m+1]) == 0 {
				r3++
			}
		}
	}
	return r1 == len(a) || r2 == len(a) || r3 == len(a)
}
