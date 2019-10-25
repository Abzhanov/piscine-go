package piscine

func IsSeparator(c rune) bool {
	return c == ' ' || c == '\n' || c == '\t'
}
func SplitWhiteSpaces(str string) []string {
	count := 0
	size := 0
	for range str {
		size++
	}
	for i, l := range str {
		if IsSeparator(l) && i > 0 && size-1 != i && !IsSeparator(rune(str[i-1])) {
			count++
		}
	}

	res := make([]string, count+1)
	j := 0
	word := ""
	for i, item := range str {
		if !IsSeparator(item) {
			word += string(item)
		}
		if IsSeparator(item) || size-1 == i && !IsSeparator(item) {
			if word != "" {
				res[j] = word
				j++
			}
			word = ""
		}
	}
	return res
}
