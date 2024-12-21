package piscine

func TrimAtoi(s string) int {
	r := 0
	neg := false
	fn := false

	for _, c := range s {
		if !fn && c == '-' {
			neg = true
		}
		if c >= '0' && c <= '9' {
			fn = true
			r = r*10 + int(c) - 48
		}
	}
	if neg {
		return -r
	}
	return r
}
