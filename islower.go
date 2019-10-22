package piscine

func IsLower(str string) bool {
	x := []rune(str)
	l := StrLen(str)
	for i := 0; i <= l-1; i++ {
		if (x[i] >= 0) && (x[i] <= 96) || (x[i] >= 123) && (x[i] <= 127) {
			return false
		}
	}
	return true
}
