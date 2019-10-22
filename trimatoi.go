package piscine

func TrimAtoi(s string) int {
	x := []rune(s)
	j := 0
	var res int
	for i := range s {
		if j == i {
			j++
		}
	}
	n := 0
	for i := 0; i < j; i++ {
		if x[i] == '+' && n == 0 {
			n = 1
		} else if x[i] == '-' && n == 0 {
			n = 2
		} else if x[i] >= 48 && s[i] <= 57 && (n == 0 || n == 3) {
			res = res*10 + (int(x[i]) - 48)
			n = 3
		} else if x[i] >= 48 && s[i] <= 57 && n == 1 {
			res = res*10 + (int(x[i]) - 48)
		} else if x[i] >= 48 && s[i] <= 57 && n == 2 {
			res = res*10 + (int(x[i]) - 48)
		}
	}
	if n == 1 || n == 3 {
		return res
	}
	return res * (-1)
}
