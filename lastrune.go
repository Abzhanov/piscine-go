package piscine

func LastRune(s string) rune {
	ar := []rune(s)
	return ar[lent(ar)-1]
}

func lent(a []rune) int {
	rev := 0
	for _, _ = range a {
		rev++
	}
	return rev
}
