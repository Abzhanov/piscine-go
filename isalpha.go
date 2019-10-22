package piscine

func IsAlpha(s string) bool {
	x := []rune(s)
	for i := 0; i <= len(x)-1; i++ {	
		if (x[i] >= 0) && (x[i] <= 47) || (x[i] >= 58) && (x[i] <= 64) || (x[i] >= 91) && (x[i] <= 96) || (x[i] >= 123) && (x[i] <= 127) {
			return false
		}
	}
	return true
}
