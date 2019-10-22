package piscine

func AlphaCount(str string) int {
	a := []rune(str)
	counter := 0
	for _, i := range a {
		if 97 <= i && i <= 122 || 65 <= i && i <= 90 {
			counter++
		}
	}
	return counter
}
