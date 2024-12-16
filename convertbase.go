package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimalValue := toDecimal(nbr, baseFrom)
	return fromDecimal(decimalValue, baseTo)
}

func toDecimal(nbr, base string) int {
	baseLen := len(base)
	decimalValue := 0

	for _, char := range nbr {
		value := indexRune(base, char)
		decimalValue = decimalValue*baseLen + value
	}

	return decimalValue
}

func fromDecimal(decimal int, base string) string {
	if decimal == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	result := ""

	for decimal > 0 {
		remainder := decimal % baseLen
		result = string(base[remainder]) + result
		decimal /= baseLen
	}

	return result
}

func indexRune(s string, r rune) int {
	for i, char := range s {
		if char == r {
			return i
		}
	}
	return -1
}
