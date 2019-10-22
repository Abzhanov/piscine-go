package piscine

func Index(s string, toFind string) int {

	str := []rune(s)
	c := []rune(toFind)
	n := len(str)
	m := len(c)

	for i := 0; i <= n-m; i++ {
		if toFind == s[i:i+m] {
			return (i)
		}
	}
	return -1
}
