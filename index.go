package piscine

func Index(s string, find string) int {
	lenstring := len(s)
	lentofind := len(find)

	if lentofind == 0 {
		return 0
	}
	for i := 0; i <= lenstring-lentofind; i++ {
		if s[i:i+lentofind] == find {
			return i
		}
	}
	return -1
}
