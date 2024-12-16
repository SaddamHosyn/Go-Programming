package piscine

func ConcatParams(args []string) string {
	output := make([]byte, 0)
	for j, arg := range args {
		output = append(output, arg...)
		if j < len(args)-1 {
			output = append(output, '\n')
		}
	}
	return string(output)
}
