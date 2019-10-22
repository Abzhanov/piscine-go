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
	last := 0
	for i := 0; i < j; i++ {
		if x[i] == '+' && last == 0 {
			last = 1
		} else if x[i] == '-' && last == 0 {
			last = 2
		} else if x[i] >= 48 && s[i] <= 57 && (last == 0 || last == 3) {
			res = res*10 + (int(x[i]) - 48)
			last = 3
		} else if x[i] >= 48 && s[i] <= 57 && last == 1 {
			res = res*10 + (int(x[i]) - 48)
		} else if x[i] >= 48 && s[i] <= 57 && last == 2 {
			res = res*10 + (int(x[i]) - 48)
		}
	}
	if last == 1 || last == 3 {
		return res
	}
	return res * (-1)
}
