package piscine

func Capitalize(s string) string {
	l := len(s)
	x := []rune(s)
	n := true
	for i := 0; i < l; i++ {
		if n && (x[i] >= 'a' && x[i] <= 'z') {
			x[i] = x[i] - 32
			n = false
		} else if n && (x[i] >= 'A' && x[i] <= 'Z') {
			n = false
		} else if n == false && (x[i] >= 'a' && x[i] <= 'z') {

		} else if n == false && (x[i] >= 'A' && x[i] <= 'Z') {
			x[i] = x[i] + 32
		} else if x[i] >= '0' && x[i] <= '9') {
			n = false
		} else {
			n = true
		}
	}
	return string(x)
}