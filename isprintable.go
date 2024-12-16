package piscine

func IsPrintable(s string) bool {
	for _, character := range s {
		if character < ' ' || character > '~' {
			return false
		}
	}
	return true
}
