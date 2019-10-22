package piscine

func IsNumeric(str string) bool {
	x := []rune(str)
	l := StrLen(str)
	for i := 0; i <= l-1; i++ {
		if (x[i] >= 0) && (x[i] <= 47) || (x[i] >= 58) && (x[i] <= 127) {
			return false
		}
	}
	return true
}
