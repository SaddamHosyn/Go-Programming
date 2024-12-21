package piscine

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	word := ""
	for i := 0; i < len(s); i++ {
		if isWhitespace(s[i]) {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
