package piscine

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	k := len(arg)
	for i := 0; i < k-1; i++ {
		for j := 0; j < k-i-1; j++ {
			if arg[j] > arg[j+1] {
				arg[j], arg[j+1] = arg[j+1], arg[j]
			}
		}
	}
	return arg[2]
}
