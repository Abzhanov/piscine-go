package piscine

func IsPrintable(str string) bool {
	x := []rune(str)
	l := StrLen(str)
	for i := 0; i <= l-1; i++ {
		if (x[i] >= 32) && (x[i] <= 127) {
			return false
		}
	}
	return true
}
