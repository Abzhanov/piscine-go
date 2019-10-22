package piscine

func NRune(s string, n int) rune {
	array := []rune(s)
	i := 0; 
	for range array {
		i++
	}	
	if n <= i && n > 0 {
		return array[n-1]
	}
	return 0
}
