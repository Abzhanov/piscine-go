package piscine

func AlphaCount(str string) int {
	a := []rune(str)
	counter := 0
	for _, i := range a {
		if 'a' <= i && i <= 'z' || 'A' <= i && i <= 'Z' {
			counter++
		}
	}
	return counter
}
