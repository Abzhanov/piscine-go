package piscine

func NRune(s string, n int) rune {
	array := []rune(s)
	for i := 0; i <= n; i++ {
		if i == n {
			return array[n-1]
		}
	}
	return 0
}
