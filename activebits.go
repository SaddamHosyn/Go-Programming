package piscine

func ActiveBits(n int) int {
	if 0 <= n && n < 2 {
		return (n)
	}
	return ((n) % 2) + ActiveBits(n/2)
}
