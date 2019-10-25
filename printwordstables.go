package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, word := range table {
		for _, c := range word {
			z01.PrintRune(rune(c))
		}
		z01.PrintRune(10)
	}
}


/*
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
*/
