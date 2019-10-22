package piscine

func ToUpper(s string) string {
	x := []rune(s)
	result := ""
	for i := 0; i <= len(x)-1; i++ {
		if (x[i] >= 'a') && (x[i] <= 'z') {
			x[i] = x[i] - 32
		}
		result += string(x[i])
	}
	return result
}
