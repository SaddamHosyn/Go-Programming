package piscine

func UltimateDivMod(d *int, f *int) {
	m := *d
	n := *f
	*d = m / n
	*f = m % n
}
