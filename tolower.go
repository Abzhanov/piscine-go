package piscine

func ToLower(s string) string {
	x := []rune(s)
	result := ""
	for i := 0; i <= len(x)-1; i++ {
		if (x[i] >= 'A') && (x[i] <= 'Z') {
			x[i] = x[i] + 32
		}
		result += string(x[i])
	}
	return result
}
