package piscine

func NRune(s string, n int) rune {
	array := []rune(s)
	for i := 0; i <= n; i++ {
		if i >= n && n > 0 {
			return (array[i-1])
		}
	}
	return 0
}
