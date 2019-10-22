package piscine

func LastRune(s string) rune {
	ar := []rune(s)
	return ar[lr(ar)-1]
}

func lr(a []rune) int {
	rev := 0
	for range a {
		rev++
	}
	return rev
}
