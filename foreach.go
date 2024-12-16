package piscine

func ForEach(f func(int), a []int) {
	for _, element := range a {
		f(element)
	}
}
