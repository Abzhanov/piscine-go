package piscine

func BasicAtoi2(s string) int {
	x := 0
	for _, y := range []rune(s) {
		if y <= '9' && y >= '0' {
			i := '0'
			j := 0
			for ; i != y; i++ {
				j++
			}
			if x != 0 || j != 0 {
				x = x*10 + j
			}
		} else {
			return 0
		}
	}
	return x
}
