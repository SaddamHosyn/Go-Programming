package piscine

func Compact(ptr *[]string) int {
	ref := *ptr
	count := 0
	for _, val := range *ptr {
		if val != "" {
			ref[count] = val
			count++
		}
	}

	*ptr = ref[0:count]

	return count
}
